CREATE TABLE estabelecimentos (
  id SERIAL PRIMARY KEY,
  nome VARCHAR(255) NOT NULL,
  numero_estabelecimento INTEGER NOT NULL,
  razao_social VARCHAR(255),
  endereco VARCHAR(150),
  cidade VARCHAR(80),
  estado VARCHAR(80),
  cep VARCHAR(15),
  numero VARCHAR(15)
);

CREATE TABLE lojas (
  id SERIAL PRIMARY KEY,
  nome VARCHAR(255) NOT NULL,
  numero_loja INTEGER NOT NULL,
  razao_social VARCHAR(255),
  endereco VARCHAR(150),
  cidade VARCHAR(80),
  estado VARCHAR(80),
  cep VARCHAR(15),
  numero VARCHAR(15),
  estabelecimento_id INTEGER NOT NULL REFERENCES estabelecimentos (id) ON DELETE RESTRICT
);