import React from 'react';
import { useProducts } from '../hooks/useProducts';
import { Card } from '../components/Card';

export const ProductPage = () => {
  const { data: products, isLoading, isError, error } = useProducts();

  if (isLoading) return (
    <div className="flex justify-center items-center h-64">
      <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-green-600"></div>
    </div>
  );
  if (isError) return <div className="p-8 text-center text-red-500 bg-red-50 rounded-lg shadow-sm border border-red-100 mx-8 my-8">Error: {(error as Error).message}</div>;

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8 animate-fade-in">
      <div className="mb-8 flex justify-between items-end">
        <div>
          <h2 className="text-3xl font-extrabold text-gray-900">Products Catalog</h2>
          <p className="text-gray-500 mt-1">Browse and manage available products.</p>
        </div>
        <button className="bg-green-600 text-white px-5 py-2.5 rounded-lg shadow hover:bg-green-700 transition font-medium">
          New Product
        </button>
      </div>

      {(!products || products.length === 0) ? (
        <div className="bg-white shadow-sm rounded-xl p-12 text-center text-gray-500 border border-gray-100">
          <svg className="mx-auto h-12 w-12 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
          </svg>
          <span className="mt-2 block text-sm font-medium text-gray-900">No products available.</span>
        </div>
      ) : (
        <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-4">
          {products.map((product) => (
            <Card 
              key={product.id} 
              title={product.name} 
              content={product.description}
              price={product.price}
            />
          ))}
        </div>
      )}
    </div>
  );
};
