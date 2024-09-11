import z from "zod";

export const loginSchema = z.object({
  email: z.string().email(),
  password: z
    .string()
    .min(1, "password required")
    .min(8, "password at least 8 characters"),
});

export const refreshTokenSchema = z.object({
  refreshToken: z.string(),
});

export const registerSchema = z
  .object({
    name: z
      .string()
      .min(1, "user name required")
      .min(3, "user name at least 3 characters"),
    email: z.string().email(),
    password: z
      .string()
      .min(1, "password required")
      .min(8, "password at least 8 characters"),
    confirmPassword: z
      .string()
      .min(1, "confirm password required")
      .min(8, "confirm password at least 8 characters"),
  })
  .refine(
    (data) => data.password === data.confirmPassword,
    "passwords should match"
  );

export type LoginInput = z.infer<typeof loginSchema>;
export type RegisterInput = z.infer<typeof registerSchema>;
export type RefreshTokenInput = z.infer<typeof refreshTokenSchema>;
