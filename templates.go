package main

const indexTemplate = `
{{define "index"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="/static/css/styles.css">
    <title>Ujay's Malunggay Crackers</title>
</head>
<body>
    <div id="background"></div>
    <div id="globe"></div>

    {{template "navbar" .}}
    {{template "home" .}}
    {{template "modals" .}}
    {{template "swot" .}}
    {{template "bmc" .}}
    {{template "sr" .}}
    {{template "contact" .}}
    {{template "footer" .}}

    <div id="toast-container" class="toast-container">
        Message sent successfully!
    </div>

    <script src="/static/js/script.js"></script>
</body>
</html>
{{end}}
`

const navbarTemplate = `
{{define "navbar"}}
<nav class="navbar">
    <div class="container">
        <div class="logo">Ujay's Malunggay Crackers</div>
        <ul class="nav-links">
            <li><a href="#home">Home</a></li>
            <li><a href="#swot">SWOT</a></li>
            <li><a href="#bmc">BMC</a></li>
            <li><a href="#sr">SR</a></li>
            <li><a href="#contact">Contact</a></li>
        </ul>
        <div class="menu-btn">☰</div>
    </div>
</nav>
{{end}}
`

const homeTemplate = `
{{define "home"}}
<section id="home">
    <div class="container home-content fade-in">
        <div class="brand-intro">
            <h1>Ujay's Malunggay Crackers</h1>
            <p>Taste the goodness of Ujay's Malunggay Crackers - a crunchy, healthy snack packed with the power of malunggay! Enjoy a delicious, guilt-free treat that's perfect for any time of the day.</p>
            <button class="btn" onclick="openModal()">
                <i class="fas fa-play"></i> Watch Our Story
            </button>
            <button class="btn" onclick="openProductModal()">
                <i class="fas fa-box-open"></i> View Products
            </button>
        </div>
        <div class="products">
            <img src="/static/img/banner-1.jpg" alt="Product 1" class="product-image active">
            <img src="/static/img/banner-2.jpg" alt="Product 2" class="product-image">
            <img src="/static/img/banner-3.jpg" alt="Product 3" class="product-image">
        </div>
    </div>
</section>
{{end}}
`

const modalsTemplate = `
{{define "modals"}}
<div id="modal" class="modal">
    <div class="modal-content">
        <span class="close" onclick="closeModal()">X</span>
        <iframe id="video-frame" width="300" height="285" frameborder="0" allowfullscreen></iframe>
    </div>
</div>

<div id="product-modal" class="modal">
    <div class="modal-content">
        <span class="close" onclick="closeProductModal()">X</span>
        <h2>Our Products</h2>
        <div class="product-gallery">
            <div class="product-item">
                <img src="/static/img/banner-1.jpg" alt="Product 1">
                <p>Malunggay Crackers - Original</p>
            </div>
            <div class="product-item">
                <img src="/static/img/banner-2.jpg" alt="Product 2">
                <p>Malunggay Crackers - Cheese</p>
            </div>
            <div class="product-item">
                <img src="/static/img/banner-3.jpg" alt="Product 3">
                <p>Malunggay Crackers - Spicy</p>
            </div>
        </div>
    </div>
</div>
{{end}}
`

const swotTemplate = `
{{define "swot"}}
<section id="swot">
    <div class="container">
        <h1 class="swot-title">SWOT Analysis</h1>
        <div class="swot-grid fade-in">
            <div class="swot-box">
                <h2 class="text-2xl">Strengths</h2>
                <ul class="list-disc">
                    <li>Stand Firm in the Healthy Product Industry</li>
                    <li>Unique and Exquisite Taste</li>
                    <li>Healthy Community of Resellers</li>
                    <li>Sustainability</li>
                    <li>Online Presence</li>
                </ul>
            </div>
            <div class="swot-box">
                <h2 class="text-2xl">Weaknesses</h2>
                <ul class="list-disc">
                    <li>New to the Market</li>
                    <li>Bootstrapping</li>
                </ul>
            </div>
            <div class="swot-box">
                <h2 class="text-2xl">Opportunities</h2>
                <ul class="list-disc">
                    <li>Create business venture</li>
                    <li>Innovation of Healthy Snacks</li>
                    <li>Taking Advantage of Social Media</li>
                </ul>
            </div>
            <div class="swot-box">
                <h2 class="text-2xl">Threats</h2>
                <ul class="list-disc">
                    <li>Competitors</li>
                    <li>Long-term Business Strategy</li>
                    <li>Price Competition</li>
                </ul>
            </div>
        </div>
    </div>
</section>
{{end}}
`

const bmcTemplate = `
{{define "bmc"}}
<section id="bmc">
    <div class="bmc-container">
        <h1 class="bmc-title">Business Model Canvas</h1>
        <div class="layout">
            <div class="sectionz partners">
                <h3 class="text-2xl">Key Partners</h3>
                <ul class="list-disc">
                    <li>Shopee (Online Supplier)</li>
                    <li>Own Grown Malunggay</li>
                </ul>
            </div>
            <div class="sectionz activities">
                <h3 class="text-2xl">Key Activities</h3>
                <ul class="list-disc">
                    <li>Business Promotion using Facebook</li>
                </ul>
            </div>
            <div class="sectionz value">
                <h3 class="text-2xl">Value Proposition</h3>
                <ul class="list-disc">
                    <li>Ensure that the snack consists of homemade malunggay crackers</li>
                </ul>
            </div>
            <div class="sectionz relationships">
                <h3 class="text-2xl">Customer Relationships</h3>
                <ul class="list-disc">
                    <li>Offers discounts for wholesale purchases.</li>
                    <li>Regular updates</li>
                </ul>
            </div>
            <div class="sectionz segments">
                <h3 class="text-2xl">Customer Segments</h3>
                <ul class="list-disc">
                    <li>Health-conscious individuals</li>
                    <li>All ages</li>
                </ul>
            </div>
            <div class="sectionz resources">
                <h3 class="text-2xl">Key Resources</h3>
                <ul class="list-disc">
                    <li>Malunggay</li>
                    <li>Packaging</li>
                    <li>Dehydrator Machine</li>
                    <li>Stickers</li>
                    <li>Sealer Machine</li>
                </ul>
            </div>
            <div class="sectionz channels">
                <h3 class="text-2xl">Channels</h3>
                <ul class="list-disc">
                    <li>Facebook</li>
                    <li>Shopee</li>
                    <li>Lazada</li>
                    <li>TikTok</li>
                </ul>
            </div>
            <div class="sectionz cost-structure">
                <h3 class="text-2xl">Cost Structure</h3>
                <ul class="list-disc">
                    <li>Utilities / Bills</li>
                    <li>Ingredients / Materials</li>
                    <li>Payroll / Labor</li>
                    <li>Transportation</li>
                </ul>
            </div>
            <div class="sectionz revenue-streams">
                <h3 class="text-2xl">Revenue Streams</h3>
                <ul class="list-disc">
                    <li>Regular Wholesalers</li>
                    <li>Direct Sales and Consignment</li>
                </ul>
            </div>
        </div>
    </div>
</section>
{{end}}
`

const srTemplate = `
{{define "sr"}}
<section id="sr">
    <div class="sr-container">
        <h1 class="sr-title">Strategic Recommendations</h1>
        <div class="sr-list">
            <div class="sr-item">
                <h2>Recommendation 1</h2>
                <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Neque doloremque dolore possimus magnam asperiores vitae incidunt nisi, accusantium culpa iure.</p>
            </div>
            <div class="sr-item">
                <h2>Recommendation 2</h2>
                <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Neque doloremque dolore possimus magnam asperiores vitae incidunt nisi, accusantium culpa iure.</p>
            </div>
            <div class="sr-item">
                <h2>Recommendation 3</h2>
                <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Neque doloremque dolore possimus magnam asperiores vitae incidunt nisi, accusantium culpa iure.</p>
            </div>
        </div>
    </div>
</section>
{{end}}
`

const contactTemplate = `
{{define "contact"}}
<section id="contact">
    <div class="contact-container">
        <h1 class="contact-title">Contact Us</h1>
        <div class="container fade-in">
            <form id="contact-form" class="contact-form">
                <input type="text" name="name" placeholder="Name" required>
                <input type="email" name="email" placeholder="Email" required>
                <textarea name="message" placeholder="Message" rows="5" required></textarea>
                <button type="submit" class="btn">Send</button>
            </form>
        </div>
    </div>
</section>
{{end}}
`

const footerTemplate = `
{{define "footer"}}
<footer>
    <div class="container footer-content">
        <div>
            <h3 class="text-xl font-bold">Ujay's Malunggay Crackers</h3>
            <p>Taste the goodness of Ujay's Malunggay Crackers - a crunchy, healthy snack packed with the power of malunggay! Enjoy a delicious, guilt-free treat that's perfect for any time of the day.</p>
        </div>
        <div>
            <h3 class="text-xl font-bold">Team</h3>
            <div class="member">
                <img src="/static/img/Mami.jpg" alt="Diana">
                <span>aquino, diana rose</span>
            </div>
            <div class="member">
                <img src="/static/img/no-img.webp" alt="David">
                <span>Agonia, david</span>
            </div>
            <div class="member">
                <img src="/static/img/no-img.webp" alt="Roji">
                <span>Del Pilar, Rhojie</span>
            </div>
            <div class="member">
                <img src="/static/img/no-img.webp" alt="Harvey">
                <span>hachero, harvey</span>
            </div>
            <div class="member">
                <img src="/static/img/no-img.webp" alt="Fred">
                <span>manaois, fred andrei</span>
            </div>
        </div>
        <div>
            <h3 class="text-xl font-bold"></h3><br><br>
            <div class="member">
                <img src="/static/img/Martillos.jpg" alt="Oli">
                <span>martillos, john nemuel</span>
            </div>
            <div class="member">
                <img src="/static/img/Gelo.jpg" alt="Gelo">
                <span>periodico, angelo</span>
            </div>
            <div class="member">
                <img src="/static/img/Samson.png" alt="Samson">
                <span>samson, bench</span>
            </div>
            <div class="member">
                <img src="/static/img/Deo.jpg" alt="Deo">
                <span>tungala, deo</span>
            </div>
        </div>
        <div>
            <h3 class="text-xl font-bold">Follow Us</h3>
            <div class="social-links">
                <a href="https://facebook.com" class="social-icon"><i class="fab fa-facebook-f"></i></a>
                <a href="https://twitter.com" class="social-icon"><i class="fab fa-twitter"></i></a>
                <a href="https://linkedin.com" class="social-icon"><i class="fab fa-linkedin-in"></i></a>
            </div>
            <br>
            <h3 class="text-xl font-bold">Contact Us</h3>
            <a class="info" href="tel:+ 0913 5761 478">+ 0913 5761 478</a><br><br>
            <h3 class="text-xl font-bold">Email Us</h3>
            <a class="info" href="mailto:ujay@gail.com">ujay@gmail.com</a>
        </div>
    </div>
    <div class="copyright">
        <p>© {{.Year}} BSIS 303 - Developed by John Oliver. All rights reserved.</p>
    </div>
</footer>
{{end}}
`
