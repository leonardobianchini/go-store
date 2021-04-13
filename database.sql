CREATE DATABASE dbteste;
\c dbteste;
CREATE TABLE IF NOT EXISTS tbUsers (
  cpf VARCHAR(16) NOT NULL PRIMARY KEY,
  private VARCHAR(16),
  incompleto VARCHAR(16),
  data_ultima_compra VARCHAR(16),
  ticket_medio VARCHAR(16),
  ticket_ultima_compra VARCHAR(16),
  loja_mais_frequente VARCHAR(18),
  loja_da_ultima_compra VARCHAR(18)
); 