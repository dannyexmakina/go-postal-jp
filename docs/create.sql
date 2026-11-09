CREATE TABLE `postal` (
  `code` varchar(7) CHARACTER SET ascii NOT NULL COMMENT '全桁',
  `first` varchar(3) CHARACTER SET ascii NOT NULL COMMENT '前半３桁',
  `last` varchar(4) CHARACTER SET ascii NOT NULL COMMENT '後半４桁',
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
  ADD KEY `code` (`code`),
  ADD KEY `prefcode` (`prefcode`),
  ADD KEY `first` (`first`),
  ADD KEY `last` (`last`),
  ADD KEY `jiscode` (`jiscode`);
