import React from "react";

export default function Login() {
  return (
    <div id="Login" className="container">
      <div className="row justify-content-center">
        <div className="col-xl-6 col-lg-6 col-md-10 d-grid gap-2">

          1<div className="form-group">
            <label htmlFor="login-email" className="form-label mt-4">
              Email address
            </label>
            <input
              type="email"
              className="form-control"
              id="login-email"
              aria-describedby="emailHelp"
              placeholder="Enter email"
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
              placeholder="Password"
            />
          </div>
          <button type="submit" className="btn btn-primary">
            Submit
          </button>
        </div>
      </div>
    </div>
  );
}
