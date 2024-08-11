ALTER TABLE notifications RENAME COLUMN author TO subject;
ALTER TABLE notification_subscriptions RENAME COLUMN author TO subject_type;
