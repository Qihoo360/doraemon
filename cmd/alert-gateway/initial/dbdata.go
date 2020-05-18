package initial

var InitialData = []string{"ALTER TABLE alert ADD UNIQUE INDEX ruleid_labels_firedat(`rule_id`, `labels`(255),`fired_at`);", `INSERT INTO  users  ( name,  password ) VALUES ('admin', 'e10adc3949ba59abbe56e057f20f883e');`}
