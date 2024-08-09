ALTER TABLE notifications RENAME COLUMN subject TO author;
ALTER TABLE notification_subscriptions RENAME COLUMN subject_type TO author;
