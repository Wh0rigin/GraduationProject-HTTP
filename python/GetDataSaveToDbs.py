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
    print(data[0])
    HOSTNAME = data[0]['datasource']['host']
    PORT = data[0]['datasource']['port']
    DATABASE = data[0]['datasource']['database']
    USERNAME = data[0]['datasource']['username']
    PASSWORD = data[0]['datasource']['password']
    CHATSET = data[0]['datasource']['charset']
    db_url = 'mysql+pymysql://{}:{}@{}:{}/{}?charset={}'.format(
        USERNAME,
        PASSWORD,
        HOSTNAME,
        PORT,
        DATABASE,
        CHATSET,
    )
    engin = create_engine(db_url)

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
    temp_time = ""
    humi_time = ""
    light_time = ""
    lamp_time = ""
    fire_time = ""
    while(True):
        Session = sessionmaker(engin)
        session = Session()
        temp = getSensor(arg[2],'temperature',json_dict['ResultObj']['AccessToken'])
        humi = getSensor(arg[2],'humidity', json_dict['ResultObj']['AccessToken'])
        lamp = getSensor(arg[2],'lamp', json_dict['ResultObj']['AccessToken'])
        light = getSensor(arg[2],'light', json_dict['ResultObj']['AccessToken'])
        fire = getSensor(arg[2],'fire', json_dict['ResultObj']['AccessToken'])
        if temp_time != temp["time"]:
            print("save...temp")
            session.add_all(
                [
                    Sensor_data(created_at=temp['time'],updated_at=temp['time'],sensor_type='temperature',value=temp['value']),
                ]
            )
            session.commit()
            temp_time = temp["time"]
        if humi_time != humi["time"]:
            print("save...humi")
            session.add_all(
                [
                    Sensor_data(created_at=humi['time'],updated_at=humi['time'],sensor_type='humidity',value=humi['value']),
                ]
            )
            session.commit()
            humi_time = humi["time"]
        if light_time != light["time"]:
            print("save...light")
            session.add_all(
                [
            
                    Sensor_data(created_at=light['time'],updated_at=light['time'],sensor_type='light',value=light['value'])
                ]
            )
            session.commit()
            light_time = light["time"]
        if lamp_time != lamp["time"]:
            print("save...lamp")
            session.add_all(
                [
                    Sensor_data(created_at=lamp['time'],updated_at=lamp['time'],sensor_type='lamp',value=lamp['value']),
                ]
            )
            session.commit()
            lamp_time = lamp["time"]
        if fire_time !=fire["time"]:
            print("save...fire")
            session.add_all(
                [
                    Sensor_data(created_at=fire['time'],updated_at=fire['time'],sensor_type='fire',value=fire['value']),
                ]
            )
            session.commit()
            fire_time = fire["time"]
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

