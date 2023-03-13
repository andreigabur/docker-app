import { useState, useEffect } from 'react';

import logo from './logo.svg';
import './App.css';

function UserRow({ user }) {
  return (
    <tr>
      <td>{user.Name}</td>
      <td>{user.Email}</td>
    </tr>
  )
}

function UsersTable({ users }) {
  const rows = [];

  users.forEach((user, key) => {
    rows.push(
      <UserRow
        user={user}
        key={key} />
    );
  });

  return (
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Email</th>
        </tr>
      </thead>
      <tbody>{rows}</tbody>
    </table>
  );
}

function App() {

  const [users, setUsers] = useState([]);

  useEffect(() => {
    fetch("http://127.0.0.1:8080/getusers", {
      method: 'GET',
    })
      .then(response => response.json())
      .then(data => {
        setUsers(data.data);
        console.log(data);
      })
  }, [])

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
      </header>
      <UsersTable users={users} />
    </div>
  );
}

export default App;
