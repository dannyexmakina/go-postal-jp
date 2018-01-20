# DB関連資料

- `update` サブコマンドまたは `convert sql` サブコマンドにより生成したSQLファイルをDBにインポートすることを想定している

## 専用スキーマ

- DBにて使用する際、下記のテーブルが存在することを前提とする

```
CREATE TABLE `postal` (
  `code` varchar(7) CHARACTER SET ascii NOT NULL COMMENT '全桁',
  `first` varchar(3) CHARACTER SET ascii NOT NULL COMMENT '前半３桁',
  `last` varchar(4) CHARACTER SET ascii NOT NULL COMMENT '後半４桁',
  `p1` varchar(1) NOT NULL COMMENT '１字目',
  `p2` varchar(1) NOT NULL COMMENT '２字目',
  `p3` varchar(1) NOT NULL COMMENT '３字目',
  `p4` varchar(1) NOT NULL COMMENT '４字目',
  `p5` varchar(1) NOT NULL COMMENT '５字目',
  `p6` varchar(1) NOT NULL COMMENT '６字目',
  `p7` varchar(1) NOT NULL COMMENT '７字目',
  `prefcode` varchar(2) CHARACTER SET ascii NOT NULL COMMENT '都道府県番号２桁',
  `pref` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '都道府県名',
  `pref_kana` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '都道府県名カナ',
  `city` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '市区郡名',
  `city_kana` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '市区郡名カナ',
  `street` varchar(192) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '町域',
  `street_kana` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '町域カナ',
  `jiscode` varchar(5) CHARACTER SET ascii NOT NULL COMMENT '全国地方公共団体コード５桁'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='郵便番号テーブル';
ALTER TABLE `postal`
  ADD PRIMARY KEY (`code`),
  ADD KEY `prefcode` (`prefcode`),
  ADD KEY `p1` (`p1`),
  ADD KEY `p2` (`p2`),
  ADD KEY `p3` (`p3`),
  ADD KEY `p4` (`p4`),
  ADD KEY `p5` (`p5`),
  ADD KEY `p6` (`p6`),
  ADD KEY `p7` (`p7`),
  ADD KEY `first` (`first`),
  ADD KEY `last` (`last`),
  ADD KEY `jiscode` (`jiscode`);
```
