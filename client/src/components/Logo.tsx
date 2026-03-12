import { Gamepad2 } from "lucide-react";

export default function () {
    return (
        <div className="flex flex-row gap-2 justify-center items-center">
            <Gamepad2 size={40} color="green" />
            <h1 className="font-extrabold text-2xl text-white">Game Vault</h1>
        </div>
    );
}
