import requests
import time
import json
import threading
from sqlalchemy import create_engine
import yaml
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

from SensorData import Sensor_data

yamlPath = '../config/application.yaml'
db_url = ''

HOSTNAME = ''
PORT = ''
DATABASE = ''
USERNAME = ''
PASSWORD = ''
CHATSET = ''

engin = None

with open(yamlPath,'rb') as f:
    data = list(yaml.safe_load_all(f))
    HOSTNAME = data['datasource']['host']
    PORT = data['datasource']['port']
    DATABASE = data['datasource']['database']
    USERNAME = data['datasource']['username']
    PASSWORD = data['datasource']['password']
    CHATSET = data['datasource']['charset']
    db_url = 'mysql+pymysql://{}:{}@{}:{}/{}?charset={}'.format(
        USERNAME,
        PASSWORD,
        HOSTNAME,
        PORT,
        DATABASE,
        CHATSET,
    )
    engin = create_engine(db_url)
Base = declarative_base(engin)

def getData(*arg):
    url = "http://api.nlecloud.com/Users/Login"
    data = {
        "Account": arg[0],
        "Password": arg[1],
        "IsRememberMe": True
    }
    ret = requests.post(url,data)
    json_dict = json.loads(ret.text)
    #print(json_dict)
    Session = sessionmaker(engin)
    session = Session()
    while(True):
        temp = getSensor(arg[2],'temperature',json_dict['ResultObj']['AccessToken'])
        humi = getSensor(arg[2],'humidity', json_dict['ResultObj']['AccessToken'])
        print("get new temperature data:"+temp['value']+",time:"+temp['time'])
        print("get new humidity data:"+humi['value']+",time:"+humi['time'])
        session.add_all(
            [
                Sensor_data(created_at=temp['time'],updated_at=temp['time'],sensor_type='temperature',value=temp['value'])
            ]
        )
        time.sleep(5)


def getSensor(deviceId:str,tag:str,token:str)->dict:
    url = "http://api.nlecloud.com/devices/"+deviceId+"/sensors/"+tag
    data = {
        "AccessToken": token
    }
    res = requests.get(url, data)
    json_dict = json.loads(res.text)
    # print(json_dict)
    return {
        'value':json_dict['ResultObj']['Value'],
        'time':json_dict['ResultObj']['RecordTime']
    }


if __name__ == '__main__':
    account = input("input your newland account:")
    password = input("input your newland password:")
    deviceId = input("input your newland deviceId:")
    print("wait...")

    thread = threading.Thread(target=getData,args=(account,password,deviceId))
    thread.start()

