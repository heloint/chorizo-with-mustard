'use client'

import { useState } from "react";
import { useForm, SubmitHandler, UseFormRegister, Path } from "react-hook-form";

interface RegisterFormInputs {
    username: string
    password: string
    password2: string
    email: string
    firstname: string
    lastname: string
    conditionsAccept: string
}

type InputProps = {
    label: string;
    type: string;
    id: string;
    fieldName: Path<RegisterFormInputs>;
    register: UseFormRegister<RegisterFormInputs>;
    placeholder: string;
    validators: Object;
}

type CheckProps = {
    label: string;
    type: string;
    id: string;
    fieldName: Path<RegisterFormInputs>;
    register: UseFormRegister<RegisterFormInputs>;
    placeholder: string;
    validators: Object;
}

type RegistrationResponse = {
    Result: string;
    ConflictError: string;
}

export default function Register() {

    const [errorMessage, setErrorMessage] = useState<string | undefined>();
    const [successMessage, setSuccessMessage] = useState<string | undefined>();
    const setConflictMessage = (conflictError: string) => {
            console.log(conflictError);
        switch (conflictError) {
            case "user_name_unique":
                setErrorMessage('Username already exists!');
                break;
            case "user_email_unique":
                setErrorMessage('Email is already used!');
                break;
            default:
                setErrorMessage(conflictError);
        }
    }

    const onSubmit: SubmitHandler<RegisterFormInputs> = async (data) => {
        console.log(data);

        const res = await fetch('http://localhost:8000/register', {
            method: "POST",
            headers: {
                "Access-Control-Allow-Origin": "*",
                "Content-Type": "text/plain"
            },
            credentials: 'include',
            body: JSON.stringify({
                username: data.username,
                password: data.password,
                email: data.email,
                firstname: data.firstname,
                lastname: data.lastname,
            })
        });

        const resData: RegistrationResponse = await res.json();
        if (!res.ok) {
            setSuccessMessage(undefined);
            switch (res.status) {
                case 409:
                    setConflictMessage(resData.ConflictError);
                    break;
            }
        } else {
            setErrorMessage(undefined);
            setSuccessMessage('Registered successfully, but the email verification must be implemented yet for production.');
        }

    };
    const { register, watch, handleSubmit, formState: { errors } } = useForm<RegisterFormInputs>();

    const Input = ({ label, type, id, fieldName, register, placeholder, validators }: InputProps) => (
        <div className="form-group">
            <label>{label}</label>
            <input id={id}
                type={type}
                className="form-control"
                placeholder={placeholder}
                {...register(fieldName, validators)} onChange={e => e.stopPropagation()} />
        </div>
    );

    return (
        <div id="Register" className="container">
            <div className="row justify-content-center my-5">
                <div className="col-xl-6 col-lg-6 col-md-10 d-grid gap-2">
                    <h3>Register</h3>
                    { successMessage && <h5 className="text-success" >{successMessage}</h5> }
                    { errorMessage && <h5 className="text-danger" >{errorMessage}</h5> }
                    <form onSubmit={handleSubmit(onSubmit)} className="d-grid gap-2">
                        {/* USERNAME */}
                        <Input id="login-username"
                            label="Username"
                            type="text"
                            fieldName="username"
                            register={register}
                            placeholder="Enter username"
                            validators={{ required: true }}
                        />
                        {errors.username && <p className="text-danger">Username is required</p>}
                        {/* END: USERNAME */}

                        {/* PASSWORD */}
                        <Input id="register-password1"
                            label="Password"
                            type="password"
                            fieldName="password"
                            register={register}
                            placeholder="Enter password"
                            validators={{
                                required: true, validate: (val: string) => {
                                    if (watch('password2') != val) {
                                        return "Passwords must match!";
                                    }
                                }
                            }}
                        />
                        {errors.password?.type === 'required' && <p className="text-danger">Password is required</p>}
                        {errors.password?.type === 'validate' && <p className="text-danger">{errors.password.message}</p>}
                        {/* END: PASSWORD */}

                        {/* REPEAT PASSWORD */}
                        <Input id="register-password2"
                            label="Repeat password"
                            type="password"
                            fieldName="password2"
                            register={register}
                            placeholder="Repeat password"
                            validators={{
                                required: true, validate: (val: string) => {
                                    if (watch('password') != val) {
                                        return "Passwords must match!";
                                    }
                                }
                            }}
                        />
                        {errors.password2?.type === 'required' && <p className="text-danger">Password confirmation is required</p>}
                        {errors.password2?.type === 'validate' && <p className="text-danger">{errors.password2.message}</p>}
                        {/* END: PASSWORD CONFIRMATION*/}

                        {/* EMAIL*/}
                        <Input id="register-email"
                            label="Email"
                            type="email"
                            fieldName="email"
                            register={register}
                            placeholder="Enter email"
                            validators={{ required: true }}
                        />
                        {errors.email && <p className="text-danger">Email is required</p>}
                        {/* END: EMAIL*/}

                        {/* FIRST NAME*/}
                        <Input id="register-firstname"
                            label="First name"
                            type="text"
                            fieldName="firstname"
                            register={register}
                            placeholder="Enter first name"
                            validators={{ required: true }}
                        />
                        {errors.firstname && <p className="text-danger">First name is required</p>}
                        {/* END: FIRST NAME*/}

                        {/* LAST NAME*/}
                        <Input id="register-lastname"
                            label="Last name"
                            type="text"
                            fieldName="lastname"
                            register={register}
                            placeholder="Enter last name"
                            validators={{ required: true }}
                        />
                        {errors.lastname && <p className="text-danger">Last name is required</p>}
                        {/* END: LAST NAME*/}

                        {/* CONDITIONS */}
                        <div className="form-group d-flex gap-2 justify-content-center">
                            <input className="form-check-input"
                                type="checkbox"
                                value=""
                                id="register-conditions"
                                {...register("conditionsAccept", { required: true })}
                            />
                            <label className="form-check-label" htmlFor="register-conditions">
                                <a href="#" target="_blank" >Accept the conditions</a>
                            </label>
                        </div>
                        {errors.conditionsAccept && <p className="text-danger text-center">Must accept the conditions.</p>}
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
