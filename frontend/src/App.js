import React, { useState, useEffect } from "react";

function App() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [users, setUsers] = useState([]);

  const fetchUsers = async () => {
    const res = await fetch("/api/users");
    const data = await res.json();
    setUsers(data);
  };

  const addUser = async () => {
    await fetch("/api/user", {
      method: "POST",
      headers: {"Content-Type": "application/json"},
      body: JSON.stringify({ name, email })
    });
    fetchUsers();
  };

  useEffect(() => { fetchUsers(); }, []);

  return (
    <div style={{ padding: "20px" }}>
      <h2>User Registration</h2>

      <input placeholder="Name" onChange={e => setName(e.target.value)} />
      <br /><br />

      <input placeholder="Email" onChange={e => setEmail(e.target.value)} />
      <br /><br />

      <button onClick={addUser}>Submit</button>

      <h3>User List:</h3>
      {users.map(u => (
        <p key={u.id}>{u.name} - {u.email}</p>
      ))}
    </div>
  );
}

export default App;
