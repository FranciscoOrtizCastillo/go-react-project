import { useEffect } from "react";
import { useState } from "react";

function App() {
  const [name, setName] = useState("");
  const [users, setUsers] = useState([])

  async function loadUsers() {
    const response = await fetch("/api/users");
    const data = await response.json();
    //console.log(data)
    setUsers(data.users)
  }

  useEffect(() => {
    loadUsers()
  }, []);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const response = await fetch("/api/users", {
      method: "POST",
      body: JSON.stringify({ name }),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const data = await response.json();
    loadUsers()
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <input
          type="name"
          placeholder="Nombre"
          onChange={(e) => setName(e.target.value)}
        />

        <button>Guardar</button>
      </form>

      <ul>
        {users && users.map(user => (
          <li key={user._id}>{user.name}</li>
        ))}
      </ul>
    </div>
  );
}

export default App;
