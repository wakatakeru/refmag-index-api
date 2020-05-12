-- Create DB
CREATE TABLE refmag.papers (
  id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  title TEXT,
  doi VARCHAR(255), 
  supplement TEXT
);

-- Grant Access
GRANT ALL ON *.* TO refmag;
