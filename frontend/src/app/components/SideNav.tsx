// import React from 'react';
import Link from 'next/link';

const SideNav = () => {
  return (
    <aside className="hidden md:block w-60 bg-gray-50 p-5 border-r border-gray-200">
      <nav>
        <ul>
          <li className="mb-4">
            <Link href="/dashboard" className="block text-lg text-blue-600 hover:underline hover:bg-blue-50 p-2 rounded">
              ダッシュボード
            </Link>
          </li>
          <li className="mb-4">
            <Link href="/todos" className="block text-lg text-blue-600 hover:underline hover:bg-blue-50 p-2 rounded">
              Todos
            </Link>
          </li>
          <li className="mb-4">
            <Link href="/sprints" className="block text-lg text-blue-600 hover:underline hover:bg-blue-50 p-2 rounded">
              Sprints
            </Link>
          </li>
        </ul>
      </nav>
    </aside>
  );
};

export default SideNav;