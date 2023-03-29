import React from "react";

export default function Register() {
    return (
        <div id="Register">
            {/* USERNAME */}
            <div className="form-group">
                <label htmlFor="register-username" className="form-label mt-4">
                    Email address
                </label>
                <input
                    type="text"
                    className="form-control"
                    id="register-username"
                    placeholder="Enter username"
                />
            </div>
            {/* END: USERNAME */}

            {/* PASSWORD */}
            <div className="form-group">
                <label htmlFor="register-password1" className="form-label mt-4">
                    Password
                </label>
                <input
                    type="password"
                    className="form-control"
                    id="register-password1"
                    placeholder="Enter password"
                />
            </div>
            {/* END: PASSWORD */}

            {/* PASSWORD CONFIRMATION */}
            <div className="form-group">
                <label htmlFor="register-password2" className="form-label mt-4">
                    Repeat password
                </label>
                <input
                    type="password"
                    className="form-control"
                    id="register-password2"
                    placeholder="Repeat password"
                />
            </div>
            {/* END: PASSWORD CONFIRMATION*/}

            {/* EMAIL*/}
            <div className="form-group">
                <label htmlFor="register-email" className="form-label mt-4">
                    Email address
                </label>
                <input
                    type="email"
                    className="form-control"
                    id="register-email"
                    placeholder="Enter email"
                />
            </div>
            {/* END: EMAIL*/}
            
        </div>
    );
}