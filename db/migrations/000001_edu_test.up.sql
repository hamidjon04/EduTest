CREATE TABLE IF NOT EXISTS students(
    id UUID PRIMARY KEY,
    student_id VARCHAR UNIQUE NOT NULL,
    name VARCHAR NOT NULL,
    lastname VARCHAR NOT NULL, 
    phone_number VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS subjects(
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS student_subjects(
    student_id UUID NOT NULL, 
    subject1 UUID NOT NULL,
    subject2 UUID NOT NULL,
    PRIMARY KEY(student_id, subject1, subject2),
    FOREIGN KEY(student_id) REFERENCES students(id) ON DELETE CASCADE,
    FOREIGN KEY(subject1) REFERENCES subjects(id) ON DELETE CASCADE,
    FOREIGN KEY(subject2) REFERENCES subjects(id) ON DELETE CASCADE,
    CHECK (subject1 <> subject2)
);

CREATE TABLE IF NOT EXISTS questions(
    id UUID PRIMARY KEY,
    subject_id UUID NOT NULL, 
    question_text VARCHAR NOT NULL,
    options JSONB NOT NULL,
    answer VARCHAR NOT NULL,
    type VARCHAR NOT NULL,
    question_image VARCHAR,
    options_image JSONB,
    answer_image VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY(subject_id) REFERENCES subjects(id) ON DELETE CASCADE   
);

CREATE TABLE IF NOT EXISTS templates(
    id UUID PRIMARY KEY,
    student_id UUID NOT NULL,
    day VARCHAR NOT NULL,
    FOREIGN KEY(student_id) REFERENCES students(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS template_questions(
    template_id UUID NOT NULL,
    question_id UUID NOT NULL,
    question_number INT NOT NULL,
    FOREIGN KEY(template_id) REFERENCES templates(id) ON DELETE CASCADE,
    FOREIGN KEY(question_id) REFERENCES questions(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS cheack_questions(
    question_id UUID NOT NULL,
    template_id UUID NOT NULL,
    answer VARCHAR,
    is_correct BOOLEAN NOT NULL,
    FOREIGN KEY(question_id) REFERENCES questions(id) ON DELETE CASCADE,
    FOREIGN KEY(template_id) REFERENCES templates(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS students_result(
    student_id UUID NOT NULL,
    template_id UUID NOT NULL,
    result JSONB NOT NULL,
    ball DECIMAL NOT NULL,
    FOREIGN KEY(student_id) REFERENCES students(id) ON DELETE CASCADE,
    FOREIGN KEY(template_id) REFERENCES templates(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS templte_answers(
    template_id UUID NOT NULL,
    answer JSONB NOT NULL,
    FOREIGN KEY(template_id) REFERENCES templates(id) ON DELETE CASCADE
);