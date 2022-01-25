import { useEffect, useState } from "react";
import { Navigate } from "react-router";

export default function IndexPage() {
    const [redirect, setRedirect] = useState("");
    const [fetched, setFetched] = useState(false);

    useEffect(() => {
        if (fetched) return;
        setFetched(true);

        if (!localStorage.token) {
            setRedirect("/login");
        } else {
            fetch("/api/tokenExists?token=" + localStorage.token)
                .then((res) => res.json())
                .then((exists) => {
                    if (exists) {
                        setRedirect("/dashboard");
                    } else {
                        setRedirect("/login");
                    }
                });
        }
    });

    return <>{redirect && <Navigate to={redirect} />}</>;
}
