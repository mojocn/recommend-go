# recommend-go golang开发的电影推荐系统

## [数据说明](http://files.grouplens.org/datasets/movielens/ml-latest-small-README.html)

## 矩阵数据结构

|      | item0 | item1 |
| :--- | :----:| ----: |
| user0| 0.6   | 0.7   |
| user1| 0.6   | 0.8   |



MovieLens数据集

（1）ratings.csv
数据格式：userId,movieId,rating,timestamp


（2）movies.csv
数据格式：movieId,title,genres

## result 结果说明

数据格式：userId, [(movieId, rating)]

userId：用户ID

movieId：电影ID

rating：推荐度