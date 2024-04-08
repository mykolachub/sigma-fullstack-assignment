import { UserTokenDTO, UserDTO } from './user';

export interface ServerResponse<T> {
  code: number;
  status: string;
  data: T;
  message?: string;
}

export interface ServerErrorResponse {
  code: number;
  status: string;
  message: string;
}

export interface LoginResponseData extends UserTokenDTO {}

export interface SignupResponseData {
  user: UserDTO;
}

export interface GetMeResponseData {
  user: UserDTO;
}

export interface GetUsersResposeData {
  users: UserDTO[];
}

export interface GetUserResponseData {
  user: UserDTO;
}

export interface CreateUserResponseData {
  user: UserDTO;
}

export interface UpdateUserResponseData {
  user: UserDTO;
}

export interface DeleteUserResponseData {
  user: UserDTO;
}
