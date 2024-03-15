export interface UserDTO {
  id: string;
  email: string;
  password: string;
  role: string;
}

export interface UserCreateDTO {
  email: string;
  password: string;
  role: string;
}
