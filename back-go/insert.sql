use study;

ALTER TABLE config ADD LetexUAT VARCHAR(255) NOT NULL COMMENT "simpletexUAT鉴权";

ALTER TABLE config ADD latex_nomal_api VARCHAR(255) NOT NULL COMMENT "simpletex普通模型API";
ALTER TABLE config ADD latex_app_id VARCHAR(255) NOT NULL COMMENT "simpletex appid";
ALTER TABLE config ADD latex_secret VARCHAR(255) NOT NULL COMMENT "simpletex latex_secret";
ALTER TABLE config ADD latex_light_api VARCHAR(255) NOT NULL COMMENT "simpletex轻量模型API";

ALTER TABLE config RENAME COLUMN LetexUAT TO latex_uat;