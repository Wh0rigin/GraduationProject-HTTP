from datetime import datetime
from sqlalchemy import Column,Float, Integer, String, DateTime, Boolean
from GetDataSaveToDbs import Base

class Sensor_data(Base):
    __tablename__ = 'sensor_data'
    id = Column(Integer, primary_key=True, autoincrement=True)
    created_at = Column(DateTime)
    updated_at = Column(DateTime)
    deleted_at = Column(DateTime)
    sensor_type = Column(String(20))
    value = Column(Float)

    _locked = Column(Boolean, default=False, nullable=False)
    def __repr__(self):
        return """
            <Sensor_data(id:%s, created_at:%s, updated_at:%s, deleted_at:%s, sensor_type:%s,value:%s)>
        """ % (self.id, self.created_at, self.updated_at, self.deleted_at, self.sensor_type,self.value)