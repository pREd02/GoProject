CREATE TABLE organizations (
 id INT PRIMARY KEY,
 name VARCHAR(255) NOT NULL,
 bin VARCHAR(255) NOT NULL,
 address VARCHAR(255) NOT NULL,
 workers VARCHAR(255) NOT NULL
);

CREATE TABLE projects (
 id INT PRIMARY KEY,
 name VARCHAR(255) NOT NULL,
 organization_id INT,
 date TIMESTAMP NOT NULL,
 owner VARCHAR(255) NOT NULL,
 FOREIGN KEY (organization_id) REFERENCES organizations(id)
);

CREATE TABLE users (
 id INT PRIMARY KEY,
 name VARCHAR(255) NOT NULL,
 Phone VARCHAR(255) NOT NULL,
 Username VARCHAR(255) NOT NULL,
 Password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE organization_users (
 organization_id INT,
 user_id INT,
 FOREIGN KEY (organization_id) REFERENCES organizations(id),
 FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE project_users (
 project_id INT,
 user_id INT,
 FOREIGN KEY (project_id) REFERENCES projects(id),
 FOREIGN KEY (user_id) REFERENCES users(id)
);
