'use client'

import { useForm, SubmitHandler, UseFormRegister, Path } from "react-hook-form";
import { useState } from "react";

interface LoginFormInputs {
    username: string
    password: string
}

type InputProps = {
    label: string;
    type: string;
    id: string;
    fieldName: Path<LoginFormInputs>;
    register: UseFormRegister<LoginFormInputs>;
    placeholder: string;
    required: boolean;
}

export default function Login() {
    const { register, handleSubmit, formState: { errors } } = useForm<LoginFormInputs>();
    const [errorMesssage, setErrorMessage] = useState<string | undefined>(undefined);

    const onSubmit: SubmitHandler<LoginFormInputs> = async (data) => {
        const res = await fetch('http://localhost:8000/login', {
            method: "POST",
            headers: {
                "Access-Control-Allow-Origin": "*",
                "Content-Type": "text/plain"
            },
            credentials: 'include',
            body: JSON.stringify({
                username: data.username,
                password: data.password
            })
        });

        if (res.ok) {
            setErrorMessage(undefined);
            window.location.href = '/';
        } else {
            setErrorMessage('Incorrect username or password.');
        }

    };

    const Input = ({ label, type, id, fieldName, register, placeholder, required }: InputProps) => (
        <>
            <div className="form-group">
                <label>{label}</label>
                <input id={id}
                    type={type}
                    className="form-control"
                    placeholder={placeholder}
                    {...register(fieldName, { required })} />
            </div>
        </>
    );

    return (
        <div id="Login" className="container">
            <div className="row justify-content-center my-5">
                <div className="col-xl-6 col-lg-6 col-md-10 d-grid gap-2">
                    <h4>Login</h4>
                    {
                        errorMesssage ?
                        <h5 className="text-danger">{errorMesssage}</h5>
                        : null
                    }
                    <form onSubmit={handleSubmit(onSubmit)} className="d-grid gap-3">
                        <Input id="login-username"
                            label="Username"
                            type="text"
                            fieldName="username"
                            register={register}
                            placeholder="Enter username"
                            required
                        />
                        {errors.username && <p className="text-danger">Username is required</p>}

                        <Input id="login-pass"
                            label="Password"
                            type="password"
                            fieldName="password"
                            register={register}
                            placeholder="Enter password"
                            required
                        />
                        {errors.password && <p className="text-danger">Username is required</p>}

                        <div className="form-group">
                            <button type="submit" className="btn btn-primary">
                                Login
                            </button>
                        </div>

                    </form>
                </div>
            </div>
        </div>
    );
}
