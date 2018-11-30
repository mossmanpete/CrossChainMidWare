# -*- coding: utf-8 -*-

import os

basedir = os.path.abspath(os.path.dirname(__file__))


class Config:
    SECRET_KEY = os.environ.get('SECRET_KEY') or 'hard to guess string'
    SQLALCHEMY_COMMIT_ON_TEARDOWN = True
    FLASKY_MAIL_SUBJECT_PREFIX = '[Flasky]'
    FLASKY_MAIL_SENDER = 'Flasky Admin <flasky@example.com>'
    FLASKY_ADMIN = os.environ.get('FLASKY_ADMIN')

    DOWNLOAD_PATH = 'download'
    ETH_SECRET_KEY = 'Q!wert123@'

    @staticmethod
    def init_app(app):
        pass


class DevelopmentConfig(Config):
    DEBUG = True

    MAIL_SERVER = 'smtp.googlemail.com'
    MAIL_PORT = 587
    MAIL_USE_TLS = True
    MAIL_USERNAME = os.environ.get('MAIL_USERNAME')
    MAIL_PASSWORD = os.environ.get('MAIL_PASSWORD')
    MONGO_HOST = 'localhost'
    MONGO_PORT = 27017
    MONGO_NAME = 'chaindb'
    MONGO_USER = 'chaindb_user'
    MONGO_PASS = 'yqr.1010'
    ETH_SECRET_KEY = 'Q!wert123@'
    ETH_URL = '127.0.0.1'
    ETH_PORT = 5544
    ETH_FEE = 0.002
    ETH_Minimum = 0.5
    VIN_SIZE = 1800
    VOUT_SIZE = 80
    BTC_HOST = 'btc_wallet'
    BTC_PORT = 60011
    BTC_COLLECT_HOST = 'localhost'
    BTC_COLLECT_PORT = 5444
    BTC_FEE =0.001
    BTC_PER_FEE = 0.00005
    ETP_PORT = 8820
    ETP_URL = 'etp_wallet'
    LTC_HOST = 'ltc_wallet'
    LTC_PORT = 60012
    LTC_COLLECT_HOST = 'localhost'
    LTC_COLLECT_PORT = 5445
    LTC_FEE =0.001
    LTC_PER_FEE = 0.00005
    HC_HOST = "hc_wallet"
    HC_PORT = 19021
    HC_COLLECT_HOST = "localhost"
    HC_COLLECT_PORT = 5447
    HC_FEE = 0.05
    HC_PER_FEE = 0.001
    # QUERY_SERVICE_HOST = "192.168.1.142"
    # QUERY_SERVICE_PORT = 5444
    SUPPORT_MIDWARE_PLUGIN_SYMBOL=["HC","ETH","BTC","LTC","ERCPAX"]


class DaConfig(Config):
    DEBUG = True
    MAIL_SERVER = 'smtp.googlemail.com'
    MAIL_PORT = 587
    MAIL_USE_TLS = True
    MAIL_USERNAME = os.environ.get('MAIL_USERNAME')
    MAIL_PASSWORD = os.environ.get('MAIL_PASSWORD')
    MONGO_USER = 'chaindb_user'
    BTC_HOST = '192.168.1.124'
    BTC_PORT = 60011
    BTC_FEE = 0.001
    MONGO_HOST = '127.0.0.1'
    MONGO_PORT = 27017
    MONGO_PASS = 'yqr.1010'
    MONGO_NAME = 'chaindb'
    HC_HOST = "127.0.0.1"
    HC_PORT = 19012
    HC_FEE = 0.001



class SunnyConfig(Config):
    DEBUG = True

    MAIL_SERVER = 'smtp.googlemail.com'
    MAIL_PORT = 587
    MAIL_USE_TLS = True
    MAIL_USERNAME = os.environ.get('MAIL_USERNAME')
    MAIL_PASSWORD = os.environ.get('MAIL_PASSWORD')
    MONGO_HOST = '192.168.1.121'
    MONGO_PORT = 27017
    MONGO_NAME = 'chaindb'
    MONGO_USER = 'chaindb_user'
    MONGO_PASS = 'yqr.1010'
    ETH_SECRET_KEY = 'Q!wert123@'
    ETH_URL = '192.168.1.121'
    ETH_PORT = 8546
    ETH_Minimum = 1
    BTC_HOST = '192.168.1.104'
    BTC_PORT = 60011


class TestingConfig(Config):
    TESTING = True
    MONGO_HOST = 'chaindb'
    MONGO_PORT = 27017
    MONGO_NAME = 'chaindb'


class ProductionConfig(Config):
    MONGO_HOST = 'chaindb'
    MONGO_PORT = 27017
    MONGO_NAME = 'chaindb'


class hzkConfig(Config) :
    DEBUG = True
    MAIL_SERVER = 'smtp.googlemail.com'
    MAIL_PORT = 587
    MAIL_USE_TLS = True
    MAIL_USERNAME = os.environ.get('MAIL_USERNAME')
    MAIL_PASSWORD = os.environ.get('MAIL_PASSWORD')
    MONGO_HOST = '192.168.1.121'
    MONGO_PORT = 27017
    MONGO_NAME = 'chaindb'
    MONGO_USER = 'chaindb_user'
    MONGO_PASS = 'yqr.1010'
    ETP_URL = '192.168.1.123'
    ETP_PORT = 8820
    ETH_Minimum = 1


config = {
    'development': DevelopmentConfig,
    'testing': TestingConfig,
    'production': ProductionConfig,
    'default': DevelopmentConfig,
    'Sunny': SunnyConfig,
    'Da': DaConfig,
    'hzk':hzkConfig
}
