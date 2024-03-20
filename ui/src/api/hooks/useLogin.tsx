import axios from 'axios';
import {useMutation, UseMutationOptions} from 'react-query';

export interface LoginResponse {
    token: string;
}

export interface LoginVariables {
    email: string;
    password: string;
}

const login = async ({email, password}: LoginVariables): Promise<string> => {
    const requestData = {
        email,
        password,
    };

    try {
        const response = await axios.post<LoginResponse>('http://localhost:8080/api/v1/login', requestData, {
            headers: {
                'Content-Type': 'application/json',
            },
        });

        return response.data.token;
    } catch (error) {
        throw new Error('Login failed');
    }
};

export const useLoginMutation = (options?: Omit<UseMutationOptions<string, Error, LoginVariables, unknown>, "mutationFn"> | undefined ) => {
    return useMutation<string, Error, LoginVariables>(login, options);
};
