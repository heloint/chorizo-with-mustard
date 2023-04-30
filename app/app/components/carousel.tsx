'use client'

import Carousel from 'react-bootstrap/Carousel';
import Image from 'next/image';
import './carousel.css';

function HomeCarousel() {
  return (
    <Carousel>
      <Carousel.Item>
        <Image
          className="d-block w-100"
          src="/images/placeholder-img.webp "
          alt="First slide"
          fill
        />
        <Carousel.Caption>
          <h3>First slide label</h3>
          <p>Nulla vitae elit libero, a pharetra augue mollis interdum.</p>
        </Carousel.Caption>
      </Carousel.Item>
      <Carousel.Item>
        <Image
          className="d-block w-100"
          src="/images/placeholder-img-2.jpg"
          alt="Second slide"
          fill
        />
        <Carousel.Caption>
          <h3>Second slide label</h3>
          <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
        </Carousel.Caption>
      </Carousel.Item>
      <Carousel.Item>
        <Image
          className="d-block w-100"
          src="/images/placeholder-img.webp"
          alt="Third slide"
          fill
        />

        <Carousel.Caption>
          <h3>Third slide label</h3>
          <p>
            Praesent commodo cursus magna, vel scelerisque nisl consectetur.
          </p>
        </Carousel.Caption>
      </Carousel.Item>
    </Carousel>
  );
}

export default HomeCarousel;
