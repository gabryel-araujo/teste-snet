import { z } from "zod";

export const estabSchema = z.object({
  nome: z.string().min(3, "Nome inválido"),
  numero_estabelecimento: z.coerce.number().nonnegative(),
  razao_social: z.string().min(3, "Razão Social inválida"),
  endereco: z.string().min(3, "Endereço inválido"),
  cidade: z.string().min(3, "Cidade inválida"),
  estado: z.string().min(3, "Estado inválida"),
  cep: z.string().min(3, "Cep inválido"),
  numero: z.string(),
});

export type EstabForm = z.infer<typeof estabSchema>;
