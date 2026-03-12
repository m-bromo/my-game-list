import type { ReactNode } from "react";

interface props {
    icon: ReactNode;
    categoryName: string;
}

export default function ({ icon, categoryName }: props) {
    return (
        <div className="flex flex-col bg-secondary font-white gap-4 border rounded-xl min-w-30 min-h-30 items-center justify-center">
            <div className="flex bg-card">{icon}</div>
            {categoryName}
        </div>
    );
}
