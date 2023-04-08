'use client'

import { SyntheticEvent, useState } from "react";


export default function Register() {

    const usernameField: HTMLInputElement = document.querySelector("#register-username") as HTMLInputElement;
    const passwordField1: HTMLInputElement = document.querySelector("#register-password1") as HTMLInputElement;
    const passwordField2: HTMLInputElement = document.querySelector("#register-password2") as HTMLInputElement;
    const emailField: HTMLInputElement = document.querySelector("#register-email") as HTMLInputElement;
    const firstNameField: HTMLInputElement = document.querySelector("#register-email") as HTMLInputElement;
    const lastNameField: HTMLInputElement = document.querySelector("#register-email") as HTMLInputElement;
    const conditionsCheckbox: HTMLInputElement = document.querySelector("#register-conditions") as HTMLInputElement;

    const [username, setUsername] = useState(usernameField?.value);
    const [password1, setPassword1] = useState(passwordField1?.value);
    const [password2, setPassword2] = useState(passwordField2?.value);
    const [email, setEmail] = useState(emailField?.value);
    const [firstName, setFirstName] = useState(firstNameField?.value);
    const [lastName, setLastName] = useState(lastNameField?.value);
    const [conditionsCheck, setConditionsCheck] = useState(conditionsCheckbox?.checked ? conditionsCheckbox.checked : false);

    const submit = (e: SyntheticEvent) => {
        e.preventDefault();
        // console.log({
        //     username,
        //     password1,
        //     password2,
        //     email,
        //     firstName,
        //     lastName,
        //     conditionsCheck
        // });


        fetch('http://localhost:8000/register', {
            method: "POST",
            headers: {
                "Access-Control-Allow-Origin": "*",
                "Content-Type": "text/plain"
            },
            credentials: 'include',
            body: JSON.stringify({
                username: username,
                password: password1,
                email: email,
                firstname: firstName,
                lastname: lastName,
            })
        });
    }


    return (
        <div id="Register" className="container">
            <div className="row justify-content-center my-5">
                <div className="col-xl-6 col-lg-6 col-md-10 d-grid gap-2">
                    <h3>Register</h3>
                    <form onSubmit={submit} className="d-grid gap-2">
                        {/* USERNAME */}
                        <div className="form-group">
                            <label htmlFor="register-username" className="form-label">
                                Username
                            </label>
                            <input
                                type="text"
                                className="form-control"
                                id="register-username"
                                placeholder="Enter username"
                                onChange={e => setUsername(e.target.value)}
                            />
                        </div>
                        {/* END: USERNAME */}

                        {/* PASSWORD */}
                        <div className="form-group">
                            <label htmlFor="register-password1" className="form-label">
                                Password
                            </label>
                            <input
                                type="password"
                                className="form-control"
                                id="register-password1"
                                placeholder="Enter password"
                                onChange={e => setPassword1(e.target.value)}
                            />
                        </div>
                        {/* END: PASSWORD */}

                        {/* PASSWORD CONFIRMATION */}
                        <div className="form-group">
                            <label htmlFor="register-password2" className="form-label">
                                Repeat password
                            </label>
                            <input
                                type="password"
                                className="form-control"
                                id="register-password2"
                                placeholder="Repeat password"
                                onChange={e => setPassword2(e.target.value)}
                            />
                        </div>
                        {/* END: PASSWORD CONFIRMATION*/}

                        {/* EMAIL*/}
                        <div className="form-group">
                            <label htmlFor="register-email" className="form-label">
                                Email address
                            </label>
                            <input
                                type="email"
                                className="form-control"
                                id="register-email"
                                placeholder="Enter email"
                                onChange={e => setEmail(e.target.value)}
                            />
                        </div>
                        {/* END: EMAIL*/}

                        {/* FIRST NAME*/}
                        <div className="form-group">
                            <label htmlFor="register-first-name" className="form-label">
                                First name
                            </label>
                            <input
                                type="text"
                                className="form-control"
                                id="register-first-name"
                                placeholder="Enter first name"
                                onChange={e => setFirstName(e.target.value)}
                            />
                        </div>
                        {/* END: FIRST NAME*/}

                        {/* LAST NAME*/}
                        <div className="form-group">
                            <label htmlFor="register-last-name" className="form-label">
                                Last name
                            </label>
                            <input
                                type="text"
                                className="form-control"
                                id="register-last-name"
                                placeholder="Enter last name"
                                onChange={e => setLastName(e.target.value)}
                            />
                        </div>
                        {/* END: LAST NAME*/}

                        {/* CONDITIONS */}
                        <div className="form-group d-flex gap-2 justify-content-center">
                            <input className="form-check-input"
                                type="checkbox"
                                value=""
                                id="register-conditions"
                                onChange={e => setConditionsCheck(e.target.checked)}
                            />
                            <label className="form-check-label" htmlFor="register-conditions">
                                <a href="#" target="_blank" >Accept the conditions</a>
                            </label>
                        </div>
                        {/* END: CONDITIONS */}


                        {/* REG. BUTTON */}
                        <div className="form-group d-flex justify-content-center">
                            <button type="submit" className="btn btn-primary btn-lg">
                                Register
                            </button>
                        </div>
                        {/* END: REG. BUTTON */}
                    </form>
                </div>
            </div>
        </div>
    );
}
