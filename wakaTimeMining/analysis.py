import pymysql
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
import Apriori
import ID3
import re
import Mongo
from apyori import apriori
from sklearn.cluster import KMeans


def read_wakatime_csv():
    waka = pd.read_csv('/home/bokket/wakatime.csv')
    print(waka)
    return waka


def select_data():
    waka = read_wakatime_csv()
    waka.isnull().sum().sort_values(ascending=False)

    waka['DATE'] = pd.to_datetime(waka['DATE'])

    clion = waka['CLion']

    golang = waka['GoLand']

    vscode = waka['VS Code']

    pycharm = waka['PyCharmCore']

    distance1 = pd.concat([clion, golang, vscode, pycharm], axis=1)
    print(distance1)

    dis1 = distance1.values.reshape(344, 4)
    tool = []
    select = [1, 2, 3, 4]
    for i in range(344):
        # print(dis[i])
        if (dis1[i][0] != 0) | (dis1[i][1] != 0) | (dis1[i][2] != 0) | (dis1[i][3] != 0):
            #         print(dis1[i])
            tool.append(list(dis1[i].astype(int)))
            # dis=dis[i]
            # print('dis[i][0] > dis[i][1]?:', dis[i][0] > dis[i][1])
    print(tool)

    for i in range(len(tool)):
        for j in range(len(tool[i])):
            if tool[i][j] != 0:
                tool[i][j] = select[j]

    print(tool)

    return tool


def rule():
    tool = select_data()
    c1 = Apriori.createC1(tool)
    d = list(map(set, tool))
    l1, supportdata0 = Apriori.scanD(d, c1, 0.5)
    # print(l1)
    # print(supportdata0)
    l, supportdata = Apriori.apriori(tool, 0)
    rules = Apriori.generateRules(l, supportdata, 0)
    print(rules)
    # print(len(rules))

    return rules


def generate():
    tool = select_data()
    s = []
    for i in tool:
        # print(i)
        s.append(list(i))
    print(s)

    for i in range(len(s)):
        print(s[i])
        k = 0
        for j in range(len(s[i])):
            if s[i][j] == 0:
                k += 1
                # print(k)
        print(k)
        if k > 2:
            s[i].append('no')
        elif k <= 2:
            s[i].append('yes')

    labels = ["Clion", "GoLand", "VS Code", "PyCharmCore"]

    calc = ID3.calcShannonEnt(s)
    print(calc)

    splitData = ID3.splitDataSet(s, 0, 1)
    print(s)
    print(' ')
    splitData = ID3.splitDataSet(s, 0, 0)
    print(splitData)

    myTree = ID3.createTree(s, labels)

    print(' ')
    print(myTree)

    return myTree


def filter():
    waka = read_wakatime_csv()
    CXX = waka[['C++']]

    C = waka[['C']]

    go = waka[['Go']]

    df = pd.concat([CXX, C, go], axis=1)

    df1 = df.values.reshape(344, 3)

    tool = []
    for i in range(344):
        # print(dis1[i])
        if (df1[i][0] != 0) | (df1[i][1] != 0) | (df1[i][2] != 0):
            tool.append(df1[i])
    print(tool)

    print(len(tool))

    SSE = []
    for k in range(1, 9):
        kmodel = KMeans(n_clusters=k)
        kmodel.fit(tool)
        SSE.append(kmodel.inertia_)  # 样本到最近的聚类中心的距离平方之和

    X = range(1, 9)
    plt.xlabel('k')
    plt.ylabel('SSE')
    plt.plot(X, SSE, 'o-')
    plt.show()

    print(kmodel.cluster_centers_)  # 查看聚类中心
    print(kmodel.labels_)  # 查看各样本对应的类别

    print(' ')
    print(SSE)

    return SSE, tool


def get_kmeans():
    _, tool = filter()
    # 取k=6进行的类别聚类
    k = 2
    # df1中读取数据并进行聚类分析
    # 调用k-means算法#
    kmodel = KMeans(n_clusters=k, n_jobs=4)  # n_jobs是并行数，一般等于CPU数较好
    kmodel.fit(tool)  # 训练模型

    cen = kmodel.cluster_centers_

    print(kmodel.cluster_centers_)  # 查看聚类中心
    print(' ')

    print(kmodel.labels_)  # 查看各样本对应的类别


# 获取连接对象
def get_mysql_connect():
    conn = pymysql.connect(host='localhost', port=3306, user='root', passwd='wxz', db='wakaTime', charset='utf8')
    return conn


def init_ID3_table():
    conn = get_mysql_connect()
    cursor = conn.cursor()  # 获取游标

    # drop= '''
    #         DROP TABLE IF EXISTS `rule`;
    # 	  '''
    #
    # sqlSetUtf8 = "ALTER TABLE `rule` CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci"
    # cursor.execute(drop)  # 初始化表
    # cursor.execute(sqlSetUtf8)  # 将表的编码格式换为utf8
    # conn.commit()

    sql = '''
                CREATE TABLE IF NOT EXISTS `ID3`(
            `id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY NOT NULL,
            `root` VARCHAR(1024) NOT NULL,
            `rchild` VARCHAR(1024) NOT NULL,
            `lchild` VARCHAR(1024) NOT NULL,
            `rvalue` INT NOT NULL,
            `lvalue` INT NOT NULL
        	);'''

    sqlSetUtf8 = "ALTER TABLE `ID3` CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci"
    cursor.execute(sql)  # 初始化表
    cursor.execute(sqlSetUtf8)  # 将表的编码格式换为utf8
    conn.commit()
    cursor.close()
    conn.close()


def init_rule_table():
    conn = get_mysql_connect()
    cursor = conn.cursor()  # 获取游标

    # drop= '''
    #         DROP TABLE IF EXISTS `rule`;
    # 	  '''
    #
    # sqlSetUtf8 = "ALTER TABLE `rule` CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci"
    # cursor.execute(drop)  # 初始化表
    # cursor.execute(sqlSetUtf8)  # 将表的编码格式换为utf8
    # conn.commit()

    sql = '''
            CREATE TABLE IF NOT EXISTS `rule`(
        `id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY NOT NULL,
        `frequentSet` VARCHAR(1024) NOT NULL,
        `consequentSet` VARCHAR(1024) NOT NULL,
        `credibility` VARCHAR(1024) 
    	);'''

    sqlSetUtf8 = "ALTER TABLE `rule` CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci"
    cursor.execute(sql)  # 初始化表
    cursor.execute(sqlSetUtf8)  # 将表的编码格式换为utf8
    conn.commit()
    cursor.close()
    conn.close()


def init_SSE_table():
    conn = get_mysql_connect()  # 获取数据库链接对象
    cursor = conn.cursor()  # 获取游标

    sql = '''
            CREATE TABLE IF NOT EXISTS `SSE`(
        `id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY NOT NULL,
        `SSE` VARCHAR(1024) NOT NULL
    	);'''

    sqlSetUtf8 = "ALTER TABLE `SSE` CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci"
    cursor.execute(sql)  # 初始化表
    cursor.execute(sqlSetUtf8)  # 将表的编码格式换为utf8
    conn.commit()
    cursor.close()
    conn.close()


def init_kmeans_table():
    conn = get_mysql_connect()  # 获取数据库链接对象
    cursor = conn.cursor()  # 获取游标

    # drop = '''
    #         DROP TABLE IF EXISTS `project_category`;
    # 	   '''
    #
    # sqlSetUtf8 = "ALTER TABLE `rule` CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci"
    # cursor.execute(drop)  # 初始化表
    # cursor.execute(sqlSetUtf8)  # 将表的编码格式换为utf8
    # conn.commit()

    sql = '''
        CREATE TABLE IF NOT EXISTS `kmeans`(
    `id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY NOT NULL,
    `HorizontalCoordinates` VARCHAR(1024) NOT NULL,
    `VerticalCoordinates` VARCHAR(1024) NOT NULL,
    `ThreeDimensionalCoordinates` VARCHAR(1024) NOT NULL
	);'''

    sqlSetUtf8 = "ALTER TABLE `kmeans` CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci"
    cursor.execute(sql)  # 初始化表
    cursor.execute(sqlSetUtf8)  # 将表的编码格式换为utf8
    conn.commit()
    cursor.close()
    conn.close()


def save_points():
    _, tool = filter()
    init_kmeans_table()

    conn = get_mysql_connect()  # 获取数据库链接对象
    cursor = conn.cursor()  # 获取游标

    #  'FrequentSet' VARCHAR(1024) NOT NULL
    #     'ConsequentSet' VARCHAR(1024) NOT NULL
    #     `Credibility` VARCHAR(1024) NOT NULL
    for i in range(len(tool)):
        sql = '''
                           INSERT INTO `kmeans`
                           (`HorizontalCoordinates`,`VerticalCoordinates`,`ThreeDimensionalCoordinates`)
                           VALUES ("%s","%s",%s);''' % (tool[i][0], tool[i][1], tool[i][2])
        print(sql, '-----sql----\n')
        cursor.execute(sql)
        conn.commit()
    cursor.close()
    conn.close()


def save_rules():
    r = rule()
    init_rule_table()

    conn = get_mysql_connect()  # 获取数据库链接对象
    cursor = conn.cursor()  # 获取游标

    #  'FrequentSet' VARCHAR(1024) NOT NULL
    #     'ConsequentSet' VARCHAR(1024) NOT NULL
    #     `Credibility` VARCHAR(1024) NOT NULL
    for i in range(len(r)):
        sql = '''
                           INSERT INTO `rule`
                           (`FrequentSet`,`ConsequentSet`,`Credibility`)
                           VALUES ("%s","%s","%s");''' % (r[i][0], r[i][1], r[i][2])
        print(sql, '-----sql----\n')
        cursor.execute(sql)
        conn.commit()
    cursor.close()
    conn.close()


# 保存info数据到Mysql
def save_SSE():
    SSElist, tool = filter()
    init_SSE_table()
    conn = get_mysql_connect()  # 获取数据库链接对象
    cursor = conn.cursor()  # 获取游标
    for i in SSElist:
        print(i)
        sql = '''
              INSERT INTO `SSE`
              (`SSE`)
              VALUES ("%s");''' % i
        # print(sql,'-----sql----\n')
        cursor.execute(sql)
        conn.commit()

    cursor.close()
    conn.close()


def save_ID3():
    tree = generate()
    t = str(tree)

    t = re.sub("\{|\'|\}}|\}}}}", "", t)

    print(t)
    r = str(t).split(',')
    print(r)

    r = ["VsCode", "Clion", "GoLand", 3, 0]
    r1 = ["GoLand", "Clion", "no", 2, 0]
    r2 = ["Clion", "yes", "GoLand", 1, 0]
    r3 = ["Clion", "yes", "no", 1, 0]
    r4 = ["GoLand", "yes", "no", 2, 0]

    result = [r, r1, r2, r3, r4]

    print(result)

    init_ID3_table()

    conn = get_mysql_connect()  # 获取数据库链接对象
    cursor = conn.cursor()  # 获取游标
    for i in range(len(result)):
        sql = '''
                                       INSERT INTO `ID3`
                                       (`root`,`rchild`,`lchild`,`rvalue`,`lvalue`)
                                       VALUES ("%s","%s","%s","%d","%d");''' % (
            result[i][0], result[i][1], result[i][2], int(result[i][3]), int(result[i][4]))
        # print(sql,'-----sql----\n')
        cursor.execute(sql)
        conn.commit()

    cursor.close()
    conn.close()


def analysis():
    # save_SSE()
    # save_rules()
    # save_points()
    # Mongo.insert_dbscan()
    save_ID3()


if __name__ == "__main__":
    # filter()
    # SSE = filter()
    # print(len(SSE))
    # for i in SSE:
    #    print(i)

    # save_SSE()
    # r = rule()
    # print(len(r))
    #
    # for i in range(len(r)):
    #     for j in range(len(r[i])):
    #         print(r[i][j])

    analysis()
