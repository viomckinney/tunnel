import { useEffect, useState } from "react";
import { Navigate } from "react-router";
import { Link } from "react-router-dom";

type Dashboard = {
    username: string;
    isAdmin: boolean;
    tunnels: string[];
    sessionCount: number;
    apiKeyCount: number;
};

export default function DashboardPage() {
    const [fetched, setFetched] = useState(false);
    const [dashboard, setDashboard] = useState(null as Dashboard | null);
    const [redirect, setRedirect] = useState("");

    useEffect(() => {
        if (fetched) return;
        setFetched(true);

        fetch("/api/dashboard", {
            headers: { Authorization: localStorage.token },
        }).then(async (res) => {
            if (res.ok) {
                setDashboard(await res.json());
            } else {
                setRedirect("/login");
            }
        });
    });

    function tunnel(tunnelName: string) {
        const tunnelDomain = tunnelName + ".t.violet.wtf";

        return (
            <li>
                <a href={"https://" + tunnelDomain}>{tunnelDomain}</a>{" "}
                <button>delete</button>
            </li>
        );
    }

    return (
        <>
            <h1>hi {dashboard?.username}!</h1>
            <div>
                {dashboard?.apiKeyCount} api keys <button>delete all</button>
            </div>
            <div>
                {dashboard?.sessionCount} sessions{" "}
                <button>log all out (including you!)</button>
            </div>
            <div>
                <h2>tunnels</h2>
                <ul>{dashboard?.tunnels.map(tunnel)}</ul>
            </div>
            {dashboard?.isAdmin && (
                <div>
                    <Link to="/admin">admin panel</Link>
                </div>
            )}
            {redirect && <Navigate to={redirect} />}
        </>
    );
}
