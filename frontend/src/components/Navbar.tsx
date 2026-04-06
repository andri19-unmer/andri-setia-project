import React from 'react';
import { Link, useLocation } from 'react-router-dom';

export const Navbar = () => {
  const location = useLocation();

  const isActive = (path: string) => {
    return location.pathname === path ? 'text-blue-600 font-bold' : 'text-gray-600 hover:text-blue-500';
  };

  return (
    <nav className="bg-white shadow-sm border-b sticky top-0 z-10">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16 items-center">
          <div className="flex-shrink-0 flex items-center">
            <h1 className="text-xl font-extrabold text-blue-600 tracking-tight">AppLogo</h1>
          </div>
          <div className="hidden sm:ml-6 sm:flex sm:space-x-8">
            <Link to="/" className={`${isActive('/')} px-3 py-2 rounded-md text-sm transition-colors`}>
              Users
            </Link>
            <Link to="/products" className={`${isActive('/products')} px-3 py-2 rounded-md text-sm transition-colors`}>
              Products
            </Link>
          </div>
        </div>
      </div>
    </nav>
  );
};
