import { Brain, Map, Swords, WandSparkles } from "lucide-react";
import type { Category } from "../models/category";
import CategoryCard from "./CategoryCard";
export default function () {
    const categories: Category[] = [
        {
            id: "1",
            name: "RPG",
            icon: <WandSparkles />,
        },
        {
            id: "2",
            name: "Ação",
            icon: <Swords />,
        },
        {
            id: "3",
            name: "Aventura",
            icon: <Map />,
        },
        {
            id: "4",
            name: "Estrategia",
            icon: <Brain />,
        },
    ];

    return (
        <div className="flex flex-col items-center gap-8">
            <h1 className="font-bold text-2xl">Categorias</h1>
            <div className="flex flex-row gap-4">
                {categories.map((category) => (
                    <CategoryCard
                        key={category.id}
                        categoryName={category.name}
                        icon={category.icon}
                    />
                ))}
            </div>
        </div>
    );
}
