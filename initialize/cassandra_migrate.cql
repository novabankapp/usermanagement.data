-- Comment
USE novabankapp;
CREATE TABLE IF NOT EXISTS user_logins (id uuid, user_id text,email text,
                                                  phone text,
                                                  first_name text,
                                                  last_name text,
                                                  user_name text,
                                                  password text,
                                                  pin text,
                                                  is_active boolean,
                                                  is_locked boolean,
                                                  created_at timestamp,
                                                  PRIMARY KEY (id)
                                                  );

CREATE TABLE IF NOT EXISTS login_otps (id uuid, user_id text,
                                               pin text,
                                               expiry_date timestamp,
                                               PRIMARY KEY (id)
    );
CREATE TABLE IF NOT EXISTS login_codes (id uuid, user_id text,
                                               code text,
                                               expiry_date timestamp,
                                               PRIMARY KEY (id)
    );
CREATE TABLE IF NOT EXISTS user_login_attempts (id uuid, user_id text,
                                                code text,
                                                ip_address text,
                                                attempts int,
                                                attempt_date timestamp,
                                                PRIMARY KEY (id)
    );
CREATE TABLE IF NOT EXISTS user_account_activities (id uuid, account_id uuid,
                                                       activity text,
                                                       ip_address text,
                                                       activity_date timestamp,
                                                       PRIMARY KEY (id)
    );

CREATE TABLE IF NOT EXISTS user_accounts (id uuid, user_id text,
                                                 is_active boolean,
                                                 is_locked boolean,
                                                 created_at timestamp,
                                                 PRIMARY KEY (id)
    );

CREATE TABLE IF NOT EXISTS kyc_compliant (id uuid,
                                                  user_id text,
                                                 has_user_details boolean,
                                                 has_residential_details boolean,
                                                 has_user_identification boolean,
                                                 has_user_income boolean,
                                                 has_user_employment boolean,
                                                 PRIMARY KEY (id)
    );


CREATE TABLE IF NOT EXISTS email_password_resets (id uuid,
                                                   user_id text,
                                                   phrase text,
                                                   expiry_date timestamp,
                                                   PRIMARY KEY (id)
    );

CREATE TABLE IF NOT EXISTS phone_password_resets (id uuid,
                                                           user_id text,
                                                           pin text,
                                                           expiry_date timestamp,
                                                           PRIMARY KEY (id)
    );
