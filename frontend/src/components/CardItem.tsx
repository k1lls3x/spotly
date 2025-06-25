// src/components/CardItem.tsx
import './CardItem.css';

interface Props {
  title: string;
  image: string;
  description: string;
  extra?: string;
}

export default function CardItem({ title, image, description, extra }: Props) {
  return (
    <div className="card-item">
      <img src={image} alt={title} className="card-image" />
      <div className="card-content">
        <h3>{title}</h3>
        <p>{description}</p>
        {extra && <small>{extra}</small>}
      </div>
    </div>
  );
}
