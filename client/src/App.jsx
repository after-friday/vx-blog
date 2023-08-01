import { useState } from "react";
import "./App.css";

function App() {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmpass, setConfirmpass] = useState("");

  const onSubmit = () => {
    if (
      username !== "" &&
      email !== "" &&
      password !== "" &&
      confirmpass !== ""
    ) {
      console.log("Ok");
      if (password === confirmpass) {
        console.log("Matched");
      } else {
        console.log("Not matched");
      }
    } else {
      console.log("Required all field");
    }
  };

  console.log({ username, email, password, confirmpass });
  return (
    <div>
      <h1>Sign Up</h1>
      <form>
        <input
          type="text"
          placeholder="Username"
          required
          onChange={(event) => {
            setUsername(event.target.value);
          }}
        />
        <input
          type="email"
          placeholder="Email"
          required
          onChange={(event) => {
            setEmail(event.target.value);
          }}
        />
        <input
          type="password"
          placeholder="Password"
          required
          onChange={(event) => {
            setPassword(event.target.value);
          }}
        />
        <input
          type="password"
          placeholder="Confirm Password"
          required
          onChange={(event) => {
            setConfirmpass(event.target.value);
          }}
        />
      </form>
      <button type="button" onClick={onSubmit}>
        Sign Up
      </button>
    </div>
  );
}

export default App;
