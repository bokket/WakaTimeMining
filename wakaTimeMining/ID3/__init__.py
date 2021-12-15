from numpy import *
from scipy import *
from math import log
import operator


def calcShannonEnt(dataSet):
    numEntries = len(dataSet)
    print(numEntries)

    # 统计每个类别出现的频数，存储在字典中
    labelCounts = {}  # 类别字典（类别的名称为键，该类别的个数为值）
    for featVec in dataSet:
        # 获取类别
        # 即列表最后一个元素
        currentLabel = featVec[-1]
        print(currentLabel)

        print(labelCounts.keys())

        if currentLabel not in labelCounts.keys():  # 还没添加到字典里的类型
            labelCounts[currentLabel] = 0;
        labelCounts[currentLabel] += 1;
        print(labelCounts[currentLabel])
        print('')
    shannonEnt = 0.0

    for key in labelCounts:  # 求出每种类型的熵
        # 待分类的可能划分在多个分类之中
        # Prob(Xi)是选择该分类的概率
        prob = float(labelCounts[key]) / numEntries  # 每种类型个数占所有的比值
        shannonEnt -= prob * log(prob, 2)
    return shannonEnt  # 返回熵


# 按照给定的特征划分数据集
# 划分数据集，度量划分数据集的熵，以便判断当前是否正确地划分了数据集，对每个特征划分数据集的结果计算一次信息熵，
# 然后判断按照哪个特征划分数据集是最好的划分方式。

def splitDataSet(dataSet, axis, value):
    retDataSet = []
    for featVec in dataSet:  # 按dataSet矩阵中的第axis列的值等于value的分数据集
        # print(featVec[axis])
        # 第一个元素
        if featVec[axis] == value:  # 值等于value的，每一行为新的列表（去除第axis个数据）
            reducedFeatVec = featVec[:axis]

            reducedFeatVec.extend(featVec[axis + 1:])
            retDataSet.append(reducedFeatVec)
            print(retDataSet)
    return retDataSet  # 返回分类后的新矩阵


# 选择最好的数据集划分方式
# 遍历整个数据集，循环计算香农熵和splitDataSet()函数，找到最好的特征划分方式
# 熵计算来划分数据集。
def chooseBestFeatureToSplit(dataSet):
    # 4
    numFeatures = len(dataSet[0]) - 1  # 求属性的个数
    baseEntropy = calcShannonEnt(dataSet)
    bestInfoGain = 0.0
    bestFeature = -1
    for i in range(numFeatures):
        # 创建唯一的分类标签列表
        featList = [example[i] for example in dataSet]

        # 去重
        uniqueVals = set(featList)
        newEntropy = 0.0
        for value in uniqueVals:
            # 求第i列属性每个不同值的熵*他们的概率
            subDataSet = splitDataSet(dataSet, i, value)
            prob = len(subDataSet) / float(len(dataSet))  # 求出该值在i列属性中的概率

            # 求i列属性各值对于的熵求和
            newEntropy += prob * calcShannonEnt(subDataSet)
            print(newEntropy)
            print(baseEntropy)

        infoGain = baseEntropy - newEntropy  # 求出第i列属性的信息增益
        print(infoGain)

        # 计算最好的信息增益
        if (infoGain > bestInfoGain):
            bestInfoGain = infoGain
            bestFeature = i

    return bestFeature


# 创建树
def createTree(dataSet, labels):
    # 最后一个元素
    classList = [example[-1] for example in dataSet];

    # 如果存在元素类别相同
    if classList.count(classList[0]) == len(classList):
        return classList[0];

    # 选择信息增益最大的属性进行分（返回值是属性类型列表的下标）
    # labels = ['age', 'job','building','credit']
    bestFeat = chooseBestFeatureToSplit(dataSet);
    print("bestFeat", bestFeat)
    # 根据下表找属性名称当树的根节点
    bestFeatLabel = labels[bestFeat]
    print(bestFeatLabel)
    myTree = {bestFeatLabel: {}}  # 以bestFeatLabel为根节点建一个空树

    del (labels[bestFeat])  # 从属性列表中删掉已经被选出来当根节点的属性
    featValues = [example[bestFeat] for example in dataSet]  # 找出该属性所有训练数据的值（创建列表）

    # 去重
    uniqueVals = set(featValues)
    for value in uniqueVals:  # 根据该属性的值求树的各个分支
        subLabels = labels[:]
        myTree[bestFeatLabel][value] = createTree(splitDataSet(dataSet, bestFeat, value), subLabels)  # 根据各个分支递归创建树
    return myTree  # 生成的树


