import React from 'react';

interface CardProps {
  title: string;
  subtitle?: string;
  content?: string;
  price?: number;
}

export const Card: React.FC<CardProps> = ({ title, subtitle, content, price }) => {
  return (
    <div className="bg-white shadow-sm hover:shadow-lg rounded-xl p-6 border border-gray-100 transition-all duration-300">
      <h3 className="text-lg font-semibold text-gray-900">{title}</h3>
      {subtitle && <p className="text-sm text-gray-500 mt-1">{subtitle}</p>}
      {content && <p className="text-gray-700 mt-4 leading-relaxed">{content}</p>}
      {price !== undefined && (
        <div className="mt-5 pt-4 border-t border-gray-50 flex items-center justify-between">
          <span className="text-xl font-bold text-blue-600">${price.toFixed(2)}</span>
        </div>
      )}
    </div>
  );
};
