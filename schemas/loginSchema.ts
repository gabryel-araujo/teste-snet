import { z } from "zod";

export const loginSchema = z.object({
  email: z.string().email("Email inválido"),
  password: z.string().min(3, "A senha precisa de no mínimo 3 caracteres"),
});

export type LoginForm = z.infer<typeof loginSchema>;
