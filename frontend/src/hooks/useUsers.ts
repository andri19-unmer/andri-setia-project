import { useQuery } from '@tanstack/react-query';
import { api } from '../services/api';
import { User } from '../types';

export const useUsers = () => {
  return useQuery({
    queryKey: ['users'],
    queryFn: async () => {
      const { data } = await api.get<User[]>('/users');
      return data;
    },
  });
};
