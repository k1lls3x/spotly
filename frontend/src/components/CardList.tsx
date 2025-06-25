// src/components/CardList.tsx
import CardItem from './CardItem';
import './CardList.css';

interface CardData {
  id: string;
  title: string;
  image: string;
  description: string;
  extra?: string;
}

interface Props {
  cards: CardData[];
}

export default function CardList({ cards }: Props) {
  return (
    <div className="card-list">
      {cards.map((card) => (
        <CardItem
          key={card.id}
          title={card.title}
          image={card.image}
          description={card.description}
          extra={card.extra}
        />
      ))}
    </div>
  );
}


