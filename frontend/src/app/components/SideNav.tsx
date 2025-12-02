import React from 'react';
import Link from 'next/link';

const SideNav = () => {
  return (
    <aside className="w-60 bg-gray-50 p-5 border-r border-gray-200">
      <nav>
        <ul>
          <li className="mb-4">
            <Link href="/" className="text-lg text-blue-600 hover:underline">
              Home
            </Link>
          </li>
          <li className="mb-4">
            <Link href="/todos" className="text-lg text-blue-600 hover:underline">
              Todos
            </Link>
          </li>
          <li className="mb-4">
            <Link href="/sprints" className="text-lg text-blue-600 hover:underline">
              Sprints
            </Link>
          </li>
        </ul>
      </nav>
    </aside>
  );
};

export default SideNav;