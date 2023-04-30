import Link from "next/link";

type User = {
    Id: number
    Role: string
    Password: string
    Email: string
    Username: string
    Firstname: string
    Lastname: string
    RegistrationDate: string
}

type SessionVar = {
    IsLoggedIn: boolean
    User: User
}

export default function Navbar(session: SessionVar) {
    return (
        <nav className="navbar navbar-expand-lg navbar-dark bg-primary">
            <div className="container-fluid">
                <Link href="/" className="navbar-brand">Home</Link>
                <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarColor01" aria-controls="navbarColor01" aria-expanded="false" aria-label="Toggle navigation">
                    <span className="navbar-toggler-icon"></span>
                </button>
                <div className="collapse navbar-collapse" id="navbarColor01">
                    <ul className="navbar-nav me-auto">
                        {
                            session.IsLoggedIn ? null :
                                <li className="nav-item">
                                    <Link href="/login" className="nav-link" >Login</Link>
                                </li>
                        }
                        <li className="nav-item">
                            <Link href="/register" className="nav-link" >Register</Link>
                        </li>
                        {
                            session.IsLoggedIn ?
                                <li className="nav-item">
                                    <Link href="/logout" className="nav-link" >Log out</Link>
                                </li> : null
                        }
                    </ul>
                </div>
            </div>
        </nav>
    );
}
