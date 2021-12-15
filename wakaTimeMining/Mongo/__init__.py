import pymongo
import json

import analysis


def get_mongo_connect():
    client = pymongo.MongoClient()
    db = client["wakaTime"]
    # name = client.database_names()
    # print(name)
    return db


def openfile():
    # 打开文本文件
    url = '/home/bokket/dbscan.txt'

    file = open(url, 'r')
    js = file.read()
    dic = json.loads(js)
    print(dic)
    file.close()
    return dic


def insert_dbscan():
    # dic = openfile()

    db = get_mongo_connect()
    # table
    cluster = db['dbscan']

    rep = cluster.insert_many([{"Point": "1",
                                "HorizontalCoordinates": [0, 1991, 0, 0, 0, 0, 0, 1326, 0, 1232, 0, 0, 0, 0, 0, 0, 0, 0,
                                                          0, 0, 0, 0, 0, 0, 0, 0, 2776, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
                                                          0, 20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
                                                          0,
                                                          0, 0, 22, 2504],
                                "VerticalCoordinates": [3520, 269, 594, 2336, 5283, 4630, 374, 0, 3800, 106, 1219, 2003,
                                                        6210, 215, 6454, 2854, 4117, 6936, 41, 4367, 5026, 1484, 6048,
                                                        8574, 1615, 172, 0, 3778, 3055, 2036, 474, 7, 65, 190, 2533,
                                                        919,
                                                        2173, 2963, 147, 0, 1236, 4955, 8314, 1593, 8988, 5266, 296,
                                                        3124,
                                                        4143, 9061, 2587, 2953, 2706, 973, 3197, 7865, 1327, 12, 1040,
                                                        851, 5781, 7019, 55]},
                               {"Point": "2",
                                "HorizontalCoordinates": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                                "VerticalCoordinates": [12491, 14384, 12349, 12079, 10607, 10500, 10847, 11039, 13482,
                                                       11688, 12775]}
                               ])

    print(rep.inserted_ids)

