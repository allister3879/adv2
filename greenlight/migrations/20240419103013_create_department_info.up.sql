CREATE TABLE IF NOT EXISTS department_info (
    id SERIAL PRIMARY KEY,
    department_name TEXT NOT NULL,
    staff_quantity INTEGER NOT NULL,
    department_director TEXT NOT NULL,
    module_id INTEGER REFERENCES module_info(id)
);
