// Mobile Menu Toggle
const menuBtn = document.querySelector('.menu-btn');
const navLinks = document.querySelector('.nav-links');

menuBtn.addEventListener('click', () => {
    navLinks.classList.toggle('active');
});

// Video Modal Functions
const videoSrc = "/static/video/UJAY'S Pitch-Entrep.mp4";

function openModal() {
    const modal = document.getElementById('modal');
    const videoFrame = document.getElementById('video-frame');

    modal.style.display = 'flex';
    setTimeout(() => modal.classList.add('active'), 10);
    videoFrame.src = videoSrc;
}

function closeModal() {
    const modal = document.getElementById('modal');
    const videoFrame = document.getElementById('video-frame');

    modal.classList.remove('active');
    setTimeout(() => {
        modal.style.display = 'none';
        videoFrame.src = "";
    }, 300);
}

// Product Modal Functions
function openProductModal() {
    const modal = document.getElementById('product-modal');
    modal.style.display = 'flex';
    setTimeout(() => modal.classList.add('active'), 10);
}

function closeProductModal() {
    const modal = document.getElementById('product-modal');
    modal.classList.remove('active');
    setTimeout(() => {
        modal.style.display = 'none';
    }, 300);
}

// Scroll Animation & Highlight
const fadeIns = document.querySelectorAll('.fade-in');
const navItems = document.querySelectorAll('.nav-links a');
const sections = document.querySelectorAll('section');

function checkVisibility() {
    fadeIns.forEach(element => {
        const rect = element.getBoundingClientRect();
        if (rect.top < window.innerHeight && rect.bottom > 0) {
            element.classList.add('visible');
        }
    });

    sections.forEach(section => {
        const rect = section.getBoundingClientRect();
        if (rect.top <= 100 && rect.bottom >= 100) {
            const id = section.getAttribute('id');
            navItems.forEach(item => {
                item.classList.remove('active');
                if (item.getAttribute('href') === `#${id}`) {
                    item.classList.add('active');
                }
            });
        }
    });
}

window.addEventListener('scroll', checkVisibility);
window.addEventListener('load', checkVisibility);

// Moving Stars, Floating Objects, and Scroll Interaction
const background = document.getElementById('background');
const globe = document.getElementById('globe');
const numStars = 150;
const numObjects = 20;

for (let i = 0; i < numStars; i++) {
    const star = document.createElement('div');
    star.className = 'star';
    const size = Math.random() * 3 + 1;
    star.style.width = `${size}px`;
    star.style.height = `${size}px`;
    star.style.left = `${Math.random() * 100}vw`;
    star.style.top = `${Math.random() * 100}vh`;
    star.style.animationDuration = `${Math.random() * 5 + 5}s`;
    star.style.setProperty('--move-x', `${(Math.random() - 0.5) * 100}px`);
    star.style.setProperty('--move-y', `${(Math.random() - 0.5) * 100}px`);
    star.dataset.depth = Math.random() * 100;
    background.appendChild(star);
}

for (let i = 0; i < numObjects; i++) {
    const obj = document.createElement('div');
    obj.className = 'float-object';
    const size = Math.random() * 20 + 10;
    obj.style.width = `${size}px`;
    obj.style.height = `${size}px`;
    obj.style.left = `${Math.random() * 100}vw`;
    obj.style.top = `${Math.random() * 100}vh`;
    obj.style.animationDuration = `${Math.random() * 8 + 8}s`;
    obj.style.setProperty('--move-x', `${(Math.random() - 0.5) * 200}px`);
    obj.style.setProperty('--move-y', `${(Math.random() - 0.5) * 200}px`);
    obj.dataset.depth = Math.random() * 200;
    background.appendChild(obj);
}

window.addEventListener('scroll', () => {
    const scrollY = window.scrollY;
    const stars = document.querySelectorAll('.star');
    const objects = document.querySelectorAll('.float-object');

    stars.forEach(star => {
        const depth = parseFloat(star.dataset.depth);
        const moveY = scrollY * (depth / 100) * 0.05;
        star.style.transform = `translateY(${-moveY}px)`;
    });

    objects.forEach(obj => {
        const depth = parseFloat(obj.dataset.depth);
        const moveY = scrollY * (depth / 200) * 0.03;
        obj.style.transform = `translateY(${-moveY}px)`;
    });

    globe.style.transform = `translate(-50%, -50%) rotateX(60deg) translateZ(${-scrollY * 0.1}px)`;
});

document.addEventListener('mousemove', (e) => {
    const mouseX = e.clientX / window.innerWidth - 0.5;
    const mouseY = e.clientY / window.innerHeight - 0.5;

    const homeContent = document.querySelector('.home-content');
    if (homeContent) {
        homeContent.style.transform = `perspective(1000px) rotateY(${mouseX * 5}deg) rotateX(${-mouseY * 5}deg)`;
    }
});

let index = 0;
const images = document.querySelectorAll(".product-image");

function showNextImage() {
    images[index].classList.remove("active");
    index = (index + 1) % images.length;
    images[index].classList.add("active");
}

setInterval(showNextImage, 3000);

// Form Submission with Toast
document.getElementById('contact-form').addEventListener('submit', function(e) {
    e.preventDefault(); // Prevent default form submission
    
    // Show toast
    const toast = document.getElementById('toast-container');
    toast.classList.add('show');
    
    // Hide toast and reset form after 5 seconds
    setTimeout(() => {
        toast.classList.remove('show');
        this.reset(); // Reset the form
    }, 5000);
});