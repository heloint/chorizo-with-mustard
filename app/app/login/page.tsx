'use client'

import { SyntheticEvent, useState } from 'react';

export default function Login() {

    const usernameField: HTMLInputElement = document.querySelector("#login-username") as HTMLInputElement;
    const passwordField: HTMLInputElement = document.querySelector("#login-username") as HTMLInputElement;

    const [username, setUsername] = useState(usernameField?.value);
    const [password, setPassword] = useState(passwordField?.value);

    const submit = (e: SyntheticEvent) => {
        e.preventDefault();

        fetch('http://localhost:8000/login', {
            method: "POST",
            headers: {
                "Access-Control-Allow-Origin": "*",
                "Content-Type": "text/plain"
            },
            credentials: 'include',
            body: JSON.stringify( {
              username: username,
              password: password
            })
        });
    }

  return (
    <div id="Login" className="container">
      <div className="row justify-content-center">
        <div className="col-xl-6 col-lg-6 col-md-10 d-grid gap-2">
            <form onSubmit={submit}>
                <div className="form-group">
                  <label htmlFor="login-username" className="form-label mt-4">
                    Username
                  </label>
                  <input
                    type="text"
                    className="form-control"
                    id="login-username"
                    name="username"
                    placeholder="Enter username"
                    onChange={e => setUsername(e.target.value)}
                  />
                </div>

                <div className="form-group">
                  <label htmlFor="login-pass" className="form-label mt-4">
                    Password
                  </label>
                  <input
                    type="password"
                    className="form-control"
                    id="login-pass"
                    name="password"
                    placeholder="Password"
                    onChange={e => setPassword(e.target.value)}
                  />
                </div>
                <button type="submit" className="btn btn-primary">
                  Login
                </button>
            </form>

        </div>
      </div>
    </div>
  );
}
