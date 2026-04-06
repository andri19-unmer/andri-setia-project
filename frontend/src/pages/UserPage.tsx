import React from 'react';
import { useUsers } from '../hooks/useUsers';
import { Card } from '../components/Card';

export const UserPage = () => {
  const { data: users, isLoading, isError, error } = useUsers();

  if (isLoading) return (
    <div className="flex justify-center items-center h-64">
      <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
    </div>
  );
  if (isError) return <div className="p-8 text-center text-red-500 bg-red-50 rounded-lg shadow-sm border border-red-100 mx-8 my-8">Error: {(error as Error).message}</div>;

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8 animate-fade-in">
      <div className="mb-8 flex justify-between items-end">
        <div>
          <h2 className="text-3xl font-extrabold text-gray-900">User Management</h2>
          <p className="text-gray-500 mt-1">Manage your team members and their roles.</p>
        </div>
        <button className="bg-blue-600 text-white px-5 py-2.5 rounded-lg shadow hover:bg-blue-700 transition font-medium">
          Add User
        </button>
      </div>
      
      {(!users || users.length === 0) ? (
        <div className="bg-white shadow-sm rounded-xl p-12 text-center text-gray-500 border border-gray-100">
          <svg className="mx-auto h-12 w-12 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
          </svg>
          <span className="mt-2 block text-sm font-medium text-gray-900">No users found.</span>
        </div>
      ) : (
        <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
          {users.map((user) => (
            <Card 
              key={user.id} 
              title={user.name} 
              subtitle={user.email} 
              content={`Created: ${user.created_at ? new Date(user.created_at).toLocaleDateString() : 'N/A'}`}
            />
          ))}
        </div>
      )}
    </div>
  );
};
