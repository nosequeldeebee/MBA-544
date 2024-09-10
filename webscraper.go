package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/mattn/go-sqlite3"
)

type Faculty struct {
	Name  string
	Title string
}

func main() {
	/* SETUP:
	create database
		sqlite3 faculty.db
		.open faculty.db
	go get dependencies above
	*/
	// Open SQLite database
	db, err := sql.Open("sqlite3", "./faculty.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS faculty (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		title TEXT
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Scrape faculty data
	faculties, err := scrapeFaculty()
	if err != nil {
		log.Fatal(err)
	}

	// Insert data into database
	for _, f := range faculties {
		_, err := db.Exec("INSERT INTO faculty (name, title) VALUES (?, ?)", f.Name, f.Title)
		if err != nil {
			log.Printf("Error inserting faculty: %v", err)
		}
	}

	fmt.Println("all done scraping and printing to the database")

}

func scrapeFaculty() ([]Faculty, error) {
	// For demonstration, we'll use a string reader instead of an actual HTTP request
	htmlContent := websiteContent
	reader := strings.NewReader(htmlContent)

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	var faculties []Faculty

	doc.Find(".list-group-item").Each(func(i int, s *goquery.Selection) {
		name := s.Find(".profile__name").Text()
		title := s.Find(".profile__title").Text()
		name = strings.TrimSpace(name)
		title = strings.TrimSpace(title)
		if name != "" && title != "" {
			faculties = append(faculties, Faculty{Name: name, Title: title})
		}
	})

	return faculties, nil
}


var websiteContent = `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" lang="en">
	<head>
		<meta charset="utf-8"/>
		<meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport"/>
		<title>Faculty - Gustavson School of Business - UVic</title>
																				<meta content="1701113799320" name="Last-Modified"/>
				<meta content="1545257999056382" property="fb:app_id"/>
				<meta content="website" property="og:type"/>
				<meta content="@uvic" name="twitter:site"/>
				<meta content="UVic.ca" property="og:site_name"/>
				<meta content=" Faculty - Gustavson - University of Victoria" name="twitter:title"/>
				<meta content=" Faculty - Gustavson - University of Victoria" property="og:title"/>
				<meta content="At Gustavson, we consider ourselves a family with a vision for who we are and where we want to go." name="description" property="og:description"/>
				<meta content="https://www.uvic.ca/assets/core-4-0/img/uvic-social-media-default.gif" name="twitter:image"/>
				<meta content="https://www.uvic.ca/assets/core-4-0/img/uvic-social-media-default.gif" property="og:image"/>
				<meta content="https://www.uvic.ca/gustavson/faculty-and-research/faculty/index.php" property="og:url"/>
				<link href="https://www.uvic.ca/gustavson/faculty-and-research/faculty/index.php" rel="canonical"/>
		<meta content="b5a27b8b8e68c44a73c580ff399a991c" name="id"/>


<link href="/assets/core-4-1/img/favicon-32.png" rel="icon"/>
<!-- Stylesheet -->


		<link href="/assets/core-4-1/css/core.1725922785.css" rel="stylesheet"/>    



		<link href="/assets/core-4-1/css/animation/aos.1690245293.css" rel="stylesheet"/>    



		<link href="/assets/core-4-1/css/animation/splide.min.1690245299.css" rel="stylesheet"/>    



		<link href="/assets/core-4-1/css/animation/plyr.1690245298.css" rel="stylesheet"/>    





		<link href="/gustavson/_assets/css/site-styles.1725917907.css" rel="stylesheet"/>    











		<script>(function(w,d,s,l,i){w[l]=w[l]||[];w[l].push({'gtm.start':
		new Date().getTime(),event:'gtm.js'});var f=d.getElementsByTagName(s)[0],
		j=d.createElement(s),dl=l!='dataLayer'?'&l='+l:'';j.async=true;j.src=
		'https://www.googletagmanager.com/gtm.js?id='+i+dl;f.parentNode.insertBefore(j,f);
		})(window,document,'script','dataLayer','GTM-PF2ZCNX9');</script>





		<script src="/assets/core-4-1/js/animation/splide.min.1690244651.js" type="text/javascript"></script>    




<!-- site: Gustavson 4.1 - last modified: August 29 2024 10:15:32. -->

	</head>
	<body class="uv-info" data-offset="50" data-spy="scroll" data-target="#uv-nav-inpage">
		<header id="uv-header">










<!-- Google Tag Manager (noscript) -->
		<noscript><iframe height="0" src="https://www.googletagmanager.com/ns.html?id=GTM-PF2ZCNX9" style="display:none;visibility:hidden" title="Google Tag Manager" width="0"></iframe></noscript>
				<!-- End Google Tag Manager (noscript) -->



						<a class="sr-only sr-only-focusable visually-hidden-focusable" href="#uv-main">Skip to main content</a>



	<div>





<div class="srch collapse" id="searchMain">
		<div class="container">
				<div class="row srch__row">
						<div class="srch__lnks col d-lg-none">
																		<a aria-label="Apply quicklink" class="srch__lnk" href="/admissions/uvic-admissions/index.php">Apply</a>
																												<a aria-label="Library quicklink" class="srch__lnk" href="https://www.uvic.ca/library/">Library</a>
																												<a aria-label="A-Z quicklink" class="srch__lnk" href="/search/a-z-list/index.php">A-Z</a>
																												<a aria-label="Find a person quicklink" class="srch__lnk" href="/search/people/index.php">Find a person</a>
																												<a aria-label="Maps quicklink" class="srch__lnk" href="/search/maps-buildings/index.php">Maps</a>
																								</div>



								<form class="srch__form col-12 col-lg" action="/gustavson/search.php" method="get">
										<div class="form-group">                
												<label class="srch__lbl" for="searchUVic">Search</label>
												<input type="search" id="searchUVic" name="q" class="form-control srch__field js-autocomplete" placeholder="What are you looking for?" data-source="//www.uvic.ca/search/a-z-list/a-z-data.json" />
												<label for="searchp" class="d-none">Start page</label>
												<input type="hidden" id="searchp" name="p" value="1">
												<button type="submit" class="btn btn-light srch__btn"><i class="fa-solid fa-search"></i> <span class="sr-only">Search</span></button>
										</div>            
								</form>


																																																																												</div>
		</div>
</div>


<div class="glbl">


		<div class="container">
				<div class="glbl__row row align-items-center">
						<div class="col-auto">
								<a class="glbl__home d-flex align-items-center" href="/index.php">
										<img alt="University of Victoria logo" class="glbl__logo" src="https://www.uvic.ca/assets/core-4-0/img/uvic-wordmark-colour.svg"/>
								</a>                
						</div>
						<div class="col-auto d-flex align-items-center">
								<a aria-controls="searchMain" aria-expanded="false" aria-label="Search button" data-bs-target="#searchMain" data-bs-toggle="collapse" href="/search/index.php" id="search-btn" role="button"><i class="fa-solid fa-magnifying-glass"></i> <span class="sr-only">Search</span></a>

								<a href="/cas/login?service=https%3A%2F%2Fwww.uvic.ca%2Ftools%2Findex.php" aria-label="Sign in to UVic button" class="btn btn-sm btn-outline-primary glbl__btn glbl__btn--unauth"><i class="fa-solid fa-key"></i> Sign in</a>                                  
								<a aria-label="Go to UVic online tools" class="btn btn-sm btn-outline-primary glbl__btn glbl__btn--auth" href="/tools/index.php"><i class="fa-solid fa-wrench"></i> <span class="glbl__btn__txt">Online tools</span></a>

								<a href="/cas/logout?url=https%3A%2F%2Fwww.uvic.ca%2Findex.php" aria-label="Sign out of UVic button" class="btn btn-sm btn-outline-primary glbl__btn glbl__btn--auth"><i class="fa-solid fa-times"></i> <span class="glbl__btn__txt">Sign out</span></a>                              
						</div>
				</div>
		</div>
</div>
</div>

							<!-- header -->
		<div class="uv-hdr uv-hdr--internal">
			<div class="uv-hdr__container container">
											<button aria-controls="uvic-nav uvic-nav-audience" aria-expanded="false" class="btn uv-nav__tgl btn--ico" data-bs-target=".navbar-collapse" data-bs-toggle="collapse"><i class="fa-solid fa-bars"></i> <span class="visually-hidden">Show menu</span></button>
																				<div class="uv-hdr__name"><a aria-label="Gustavson School of Business home link" href="../../index.php"><img alt="Gustavson School of Business logo" class="py-2" src="../../_assets/images/gustavson-wordmark-white.png" style="max-height: 80px;"/></a></div>
									</div>
		</div>

															<!-- Navigation -->
<nav aria-label="primary navigation" class="navbar navbar-expand-lg uv-nav uv-nav--topic" id="uvic-nav">
		<div class="container uv-nav__container">
				<div class="collapse navbar-collapse">
						<ul class="navbar-nav me-auto">
																																					<li class="nav-item dropdown">
																																																																																																																																																																																																																																																																																																																																																																																		<a aria-expanded="false" aria-label="Open Programs menu" class="nav-link uv-nav__item dropdown-toggle icoa icoa--down" data-bs-toggle="dropdown" href="#" role="button">Programs</a>
														<div class="dropdown-menu uv-nav__dropdown">

																																																								<a aria-label="Go to Program overview page" class="dropdown-item" href="../../programs/program-overview/index.php">Program overview</a>

																																																								<a aria-label="Go to Undergraduate page" class="dropdown-item" href="../../programs/undergraduate/index.php">Undergraduate</a>

																																																								<a aria-label="Go to Graduate page" class="dropdown-item" href="../../programs/graduate/index.php">Graduate</a>

																																																								<a aria-label="Go to Executive &amp; professional page" class="dropdown-item" href="../../programs/executive-and-professional/index.php">Executive &amp; professional</a>
																																																																		</div>
												</li>
																																																							<li class="nav-item dropdown">
																																																																																																																																																																																																																																																												<a aria-expanded="false" aria-label="Open Faculty &amp; research menu" class="nav-link uv-nav__item dropdown-toggle active icoa icoa--down" data-bs-toggle="dropdown" href="#" role="button">Faculty &amp; research</a>
														<div class="dropdown-menu uv-nav__dropdown">

																																																								<a aria-label="Go to Research at Gustavson page" class="dropdown-item" href="../research-at-gustavson/index.php">Research at Gustavson</a>

																																																								<a aria-label="Go to Our faculty page" class="dropdown-item active" href="index.php">Our faculty</a>

																																																								<a aria-label="Go to Centres &amp; collaborations page" class="dropdown-item" href="../centres-and-collaborations/index.php">Centres &amp; collaborations</a>
																																																																		</div>
												</li>
																																																							<li class="nav-item dropdown">
																																																																																																																																																																																																														<a aria-expanded="false" aria-label="Open Alumni &amp; giving menu" class="nav-link uv-nav__item dropdown-toggle icoa icoa--down" data-bs-toggle="dropdown" href="#" role="button">Alumni &amp; giving</a>
														<div class="dropdown-menu uv-nav__dropdown">

																																																								<a aria-label="Go to Alumni page" class="dropdown-item" href="../../alumni-and-giving/alumni/index.php">Alumni</a>

																																																								<a aria-label="Go to Giving to Gustavson page" class="dropdown-item" href="../../alumni-and-giving/giving-to-gustavson/index.php">Giving to Gustavson</a>
																																																																		</div>
												</li>
																																																							<li class="nav-item dropdown">
																																																																																																																																																																																																																																																																																																																																																																																																																																																																				<a aria-expanded="false" aria-label="Open About menu" class="nav-link uv-nav__item dropdown-toggle icoa icoa--down" data-bs-toggle="dropdown" href="#" role="button">About</a>
														<div class="dropdown-menu uv-nav__dropdown">

																																																								<a aria-label="Go to Strategy &amp; values page" class="dropdown-item" href="../../about/strategy-values/index.php">Strategy &amp; values</a>

																																																								<a aria-label="Go to Rankings &amp; reputation page" class="dropdown-item" href="../../about/rankings-and-reputation/index.php">Rankings &amp; reputation</a>

																																																								<a aria-label="Go to News &amp; stories page" class="dropdown-item" href="../../about/news-and-stories/index.php">News &amp; stories</a>

																																																								<a aria-label="Go to Events page" class="dropdown-item" href="../../about/events/index.php">Events</a>

																																																								<a aria-label="Go to Contact us page" class="dropdown-item" href="../../about/contact-us/index.php">Contact us</a>
																																																																		</div>
												</li>
																															</ul>
				</div>
		</div>
</nav>

									<!-- Navigation -->
<nav aria-label="audience navigation" class="navbar navbar-expand-lg uv-nav uv-nav--audience" id="uvic-nav-audience">
		<div class="container uv-nav__container">
				<div class="collapse navbar-collapse">
						<ul class="navbar-nav">
							<li class="nav-item dropdown">
																																																																																																																																												<a aria-expanded="false" aria-haspopup="true" aria-label="Open Info for menu" class="nav-link uv-nav__item dropdown-toggle d-lg-none icoa icoa--down" data-bs-toggle="dropdown" href="#" role="button"><i class="fa-solid fa-user"></i> Info for...</a>
										<div class="dropdown-menu uv-nav__dropdown uv-nav--audience__dropdown">
												<i class="fa-solid fa-user"></i>
												<span class="info-for-label d-none d-lg-block"> Info for... </span>
																																																																																																																																																																													<a aria-label="Go to Community &amp; industry page" class="dropdown-item" href="../../info-for/industry-and-community/index.php">Community &amp; industry</a>
																																																																																																																																																																																																																<a aria-label="Go to International partners page" class="dropdown-item" href="../../info-for/international-partners/index.php">International partners</a>
																																																																																																																																																																																																																<a aria-label="Go to International students page" class="dropdown-item" href="../../info-for/international-students/index.php">International students</a>
																																																																																																																																																																																																																<a aria-label="Go to Current students page" class="dropdown-item" href="../../info-for/current-students/index.php">Current students</a>
																																																																																																																																																																																																																<a aria-label="Go to Faculty &amp; staff page" class="dropdown-item" href="../../info-for/faculty-and-staff/index.php">Faculty &amp; staff</a>
																																																								</div>
								</li>
						</ul>
				</div>
		</div>
</nav>

		</header>
		<main class="uv-main uv-main--info" id="uv-main">
			<div class="container">
				<div class="row uv-main__row--head">
					<div class="uv-content">
																																																										<nav aria-label="breadcrumb">
				<ol class="breadcrumb">
																																																														<li class="breadcrumb-item"><a aria-label="Go to home page" href="../../index.php">Home</a></li> 
																																																								<li class="breadcrumb-item">
																						Faculty &amp; research
																		</li>
																																																</ol>
		</nav>



		<div class="row uv-sect uv-sect--head">
				<div class="col-12">
						<h1 id="ipn-our-faculty">Our faculty</h1>
				</div>
		</div>

					</div>
				</div>
			</div>
			<div class="container">
				<div class="row uv-main__row">
					<div class="uv-left">
																<nav aria-label="Page contents" class="uv-inpage list-group" id="uv-nav-inpage">
					<a aria-label="Go to Our faculty on this page" class="list-group-item list-group-item-action uv-inpage__first" href="#ipn-our-faculty">Our faculty</a>
																								<a aria-label="Go to Dean &amp; associate deans on this page" class="list-group-item list-group-item-action" href="#ipn-dean-associate-deans">Dean &amp; associate deans</a>
																							<a aria-label="Go to Faculty members on this page" class="list-group-item list-group-item-action" href="#ipn-faculty-members">Faculty members</a>
					</nav>

					</div>
					<div class="uv-content">










































				<div class="row uv-sect bg--linear-top">
						<div class="col-12">

		<div class="btn-row">
																																																									<a class="btn btn-primary icoa icoa--out" href="https://www.uvic.ca/search/departments/index.php?cw_inChannelLink=1&amp;qtype=dept&amp;deptq=gustavson&amp;deptid=1020&amp;">Staff directory listing</a>

																																																									<a class="btn btn-primary" href="../../about/contact-us/index.php">Gustavson contact info</a>

								</div>
																				</div>
				</div>

				<div class="row uv-sect bg--dark-blue text--white">
						<div class="col-12">
																																																																																				 <p>We consider ourselves a family with a vision for who we are and where we want to go. In all our dealings, we aspire to be open, fair, engaged and passionate in everything we do.</p>
																												</div>
				</div>

				<div class="row uv-sect bg--linear-top">
						<div class="col-12">
																		<h2 id="ipn-dean-associate-deans">Dean &amp; associate deans</h2>
																																																																																																																																																																																																								<div class="profile profile--long">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson dean Anita Bhappu" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/anita-bhappu-2024.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/bhappu-anita.php"> Anita Bhappu</a>
										</p>
										<p>
																		<strong class="profile__title">Professor; Dean</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 254</span>
																										<a class="profile__email" href="mailto:bizdean@uvic.ca">bizdean@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-721-6420">250-721-6420</a>
														</p>
																		</div>
				</div>
						</div>
																																																																																								<div class="profile profile--long">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson associate dean Graham Brown" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/graham-brown.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/brown-graham.php"> Graham Brown</a>
										</p>
										<p>
																		<strong class="profile__title">Professor; Tim Price Entrepreneurship Fellow; Associate Dean, Teaching &amp; Learning</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 250</span>
																										<a class="profile__email" href="mailto:grbrown@uvic.ca">grbrown@uvic.ca</a>
																						</p>
																		</div>
				</div>
						</div>
																																																																																								<div class="profile profile--long">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson's Vivien Corwin" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/vivien-corwin.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/corwin-vivien.php"> Vivien Corwin</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Teaching Professor; Associate Dean, People &amp; Culture; Academic Director, Master in Management Program</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 260</span>
																										<a class="profile__email" href="mailto:vcorwin@uvic.ca">vcorwin@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-721-6311">250-721-6311</a>
														</p>
																		</div>
				</div>
						</div>
																																																																																								<div class="profile profile--long">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson associate dean Mia Maki" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/mia-maki.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/maki-mia.php"> Mia Maki</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Teaching Professor; Associate Dean, External &amp; Outreach</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 414</span>
																										<a class="profile__email" href="mailto:mmaki@uvic.ca">mmaki@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-472-4513">250-472-4513</a>
														</p>
																		</div>
				</div>
						</div>
																																																																																								<div class="profile profile--long">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher and associate dean Roy Suddaby" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/roy-suddaby.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/suddaby-roy.php"> Roy Suddaby</a>
										</p>
										<p>
																		<strong class="profile__title">Professor; Associate Dean, Research &amp; Innovation; Francis G. Winspear Chair</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 472</span>
																										<a class="profile__email" href="mailto:rsuddaby@uvic.ca">rsuddaby@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-721-6401">250-721-6401</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										</div>
				</div>

				<div class="row uv-sect bg--light-blue ">
						<div class="col-12">
																																																																																																		<div class="filter" data-display="#profilesDisplay"><form class="filter__control">
<div class="form-group">
<h3><label for="fs-search">Search faculty</label></h3>
<input class="filter__search form-control" id="fs-search" name="fs-search" placeholder="Keyword search filter" type="search"/>
<div class="invalid-feedback">Your search didn't match any results, try again.</div>
</div>
</form></div>
																																																						</div>
				</div>

				<div class="row uv-sect bg--linear-top">
						<div class="col-12">
																		<h2 id="ipn-faculty-members">Faculty members</h2>


																																										<div class="filter__display" id="profilesDisplay">
				<div class="list-group filtered__parent">
						<div class="filter__error alert alert-danger collapse" role="alert"><strong>No results found.</strong> Please try another filter option.</div>        
																																																																<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson's Jen Baggs" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/jen-baggs.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/baggs-jen.php"> Jen Baggs</a>
										</p>
										<p>
																		<strong class="profile__title">Professor; Academic Director, Undergraduate Programs</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 248</span>
																										<a class="profile__email" href="mailto:jenbaggs@uvic.ca">jenbaggs@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-472-4617">250-472-4617</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Francois Bastien" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/francois-bastien.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/bastien-francois.php"> François Bastien</a>
										</p>
										<p>
																		<strong class="profile__title">Assistant Professor; Academic Director, MBA Advancing Reconciliation</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 474</span>
																										<a class="profile__email" href="mailto:fbastien@uvic.ca">fbastien@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-721-6071">250-721-6071</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Kristin Brandl" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/kristin-brandl.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/brandl-kristin.php"> Kristin Brandl</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 436</span>
																										<a class="profile__email" href="mailto:kbrandl@uvic.ca">kbrandl@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-472-4445">250-472-4445</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Generic person image" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/image-missing.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/bridge-mark.php"> Mark Bridge</a>
										</p>
										<p>
																		<strong class="profile__title">Assistant Teaching Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 416</span>
																										<a class="profile__email" href="mailto:mbridge@uvic.ca">mbridge@uvic.ca</a>
																						</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Ravee Chittoor" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/ravee-chittoor.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/chittoor-raveendra.php"> Raveendra Chittoor</a>
										</p>
										<p>
																		<strong class="profile__title">Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 426</span>
																										<a class="profile__email" href="mailto:raveec@uvic.ca">raveec@uvic.ca</a>
																						</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Mark Colgate of the Gustavson School of Business" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/mark-colgate.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/colgate-mark.php"> Mark Colgate</a>
										</p>
										<p>
																		<strong class="profile__title">Professor; Academic Director, Custom MBA program</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 446</span>
																										<a class="profile__email" href="mailto:colgate@uvic.ca">colgate@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-853-3873">250-853-3873</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Diego Coraiola" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/diego-coraiola.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/coraiola-diego.php"> Diego Coraiola</a>
										</p>
										<p>
																		<strong class="profile__title">Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 470</span>
																										<a class="profile__email" href="mailto:dcoraiola@uvic.ca">dcoraiola@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-472-5649">250-472-5649</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Rick Cotton" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/rick-cotton.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/cotton-rick.php"> Rick Cotton</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 242</span>
																										<a class="profile__email" href="mailto:rcotton@uvic.ca">rcotton@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-721-8052">250-721-8052</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Wade Danis" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/wade-danis.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/danis-wade.php"> Wade Danis</a>
										</p>
										<p>
																		<strong class="profile__title">Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 230</span>
																										<a class="profile__email" href="mailto:wdanis@uvic.ca">wdanis@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-853-3872">250-853-3872</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Qianqian Du" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/qianqian-du.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/du-qianqian.php"> Qianqian Du</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 212</span>
																										<a class="profile__email" href="mailto:qianqiandu@uvic.ca">qianqiandu@uvic.ca</a>
																						</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Sara Elias" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/sara-rsta-elias.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/elias-sara.php"> Sara Elias</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 228</span>
																										<a class="profile__email" href="mailto:selias@uvic.ca">selias@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-721-6244">250-721-6244</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Takahiro Endo" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/takahiro-endo.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/endo-takahiro.php"> Takahiro Endo</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor; CAPI Jarislowsky Chair</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 434</span>
																										<a class="profile__email" href="mailto:endot@uvic.ca">endot@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-472-5207">250-472-5207</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Stacey Fitzsimmons" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/stacey-fitzsimmons.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/fitzsimmons-stacey.php"> Stacey Fitzsimmons</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 410</span>
																										<a class="profile__email" href="mailto:sfitzsim@uvic.ca">sfitzsim@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-472-4787">250-472-4787</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Ricardo Flores" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/flores-ricardo-2024.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/flores-ricardo.php"> Ricardo Flores</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 408</span>
																										<a class="profile__email" href="mailto:ricardoflores@uvic.ca">ricardoflores@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-472-4810">250-472-4810</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Generic person image" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/image-missing.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/galang-maria-carmen.php"> Maria Carmen Galang</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 218</span>
																										<a class="profile__email" href="mailto:cgalang@uvic.ca">cgalang@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-721-6111">250-721-6111</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Dale Ganley of the Gustavson School of Business" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/dale-ganley.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/ganley-dale.php"> Dale Ganley</a>
										</p>
										<p>
																		<strong class="profile__title">Assistant Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 442</span>
																										<a class="profile__email" href="mailto:dganley@uvic.ca">dganley@uvic.ca</a>
																						</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Huachao Gao" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/huachao-gao.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/gao-huachao.php"> Huachao Gao</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor</strong>
																						</p>
																		<p>
																										<a class="profile__email" href="mailto:hcgao@uvic.ca">hcgao@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-472-5503">250-472-5503</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Adel Guitouni" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/adel-guitouni.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/guitouni-adel.php"> Adel Guitouni</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 430</span>
																										<a class="profile__email" href="mailto:adelg@uvic.ca">adelg@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-721-6428">250-721-6428</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Kerstin Heilgenberg of the Gustavson School of Business" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/kerstin-heilgenberg.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/heilgenberg-kerstin.php"> Kerstin Heilgenberg</a>
										</p>
										<p>
																		<strong class="profile__title">Assistant Teaching Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 420</span>
																										<a class="profile__email" href="mailto:kerstinh@uvic.ca">kerstinh@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-853-3573">250-853-3573</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Aloysius Marcus Kahindi" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/aloysius-marcus-kahindi.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/kahindi-aloysius-marcus.php"> Aloysius Marcus Kahindi</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor; Canada Research Chair in International Sustainable Development</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 240</span>
																										<a class="profile__email" href="mailto:kahindimarcus@uvic.ca">kahindimarcus@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-472-5968">250-472-5968</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Kimball Ketsa of the Gustavson School of Business" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/kimball-ketsa.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/ketsa-kimball.php"> Kimball Ketsa</a>
										</p>
										<p>
																		<strong class="profile__title">Assistant Teaching Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 226</span>
																										<a class="profile__email" href="mailto:kketsa@uvic.ca">kketsa@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-853-3870">250-853-3870</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Jan Kietzmann" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/jan-kietzmann.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/kietzmann-jan.php"> Jan Kietzmann</a>
										</p>
										<p>
																		<strong class="profile__title">Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 214</span>
																										<a class="profile__email" href="mailto:jkietzma@uvic.ca">jkietzma@uvic.ca</a>
																						</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson faculty Michael King" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/michael-king2024.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/king-michael.php"> Michael King</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor; Lansdowne Chair in Finance</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 246</span>
																										<a class="profile__email" href="mailto:michaelking@uvic.ca">michaelking@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-721-6425">250-721-6425</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson's Saul Klein" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/saul-klein.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/klein-saul.php"> Saul Klein</a>
										</p>
										<p>
																		<strong class="profile__title">Professor</strong>
																						</p>
																		<p>
																										<a class="profile__email" href="mailto:sklein@uvic.ca">sklein@uvic.ca</a>
																						</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson faculty member Brian Leacock" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/brian-leacock.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/leacock-brian.php"> Brian Leacock</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Teaching Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 272</span>
																										<a class="profile__email" href="mailto:bleacock@uvic.ca">bleacock@uvic.ca</a>
																						</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Andie Lee" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/andie-lee.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/lee-andie.php"> Andie Lee</a>
										</p>
										<p>
																		<strong class="profile__title">Assistant Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 202</span>
																										<a class="profile__email" href="mailto:jwalee@uvic.ca">jwalee@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-472-4926">250-472-4926</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson faculty member Brent Mainprize" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/brent-mainprize.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/mainprize-brent.php"> Brent Mainprize</a>
										</p>
										<p>
																		<strong class="profile__title">Teaching Professor; Faculty Champion (Business) for National Consortium for Indigenous Economic Development; Director of Indigenous Programs, Executive Programs</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 224</span>
																										<a class="profile__email" href="mailto:brentm@uvic.ca">brentm@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-588-1172">250-588-1172</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Basma Majerbi" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/basma-majerbi.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/majerbi-basma.php"> Basma Majerbi</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 244</span>
																										<a class="profile__email" href="mailto:majerbi@uvic.ca">majerbi@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-472-4281">250-472-4281</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson faculty member Cheryl Mitchell" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/cheryl-mitchell.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/mitchell-cheryl.php"> Cheryl Mitchell</a>
										</p>
										<p>
																		<strong class="profile__title">Assistant Teaching Professor; Academic Director, MBA in Sustainable Innovation</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 258</span>
																										<a class="profile__email" href="mailto:clmitch@uvic.ca">clmitch@uvic.ca</a>
																						</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Matt Murphy" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/matt-murphy.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/murphy-matt.php"> Matt Murphy</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 440</span>
																										<a class="profile__email" href="mailto:mmurph@uvic.ca">mmurph@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-853-3868">250-853-3868</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Sudhir Nair" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/sudhir-nair.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/nair-sudhir.php"> Sudhir Nair</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor; Director, PhD program</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 431</span>
																										<a class="profile__email" href="mailto:sudhirn@uvic.ca">sudhirn@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-721-6414">250-721-6414</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Sang Nam of the Gustavson School of Business" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/sang-nam.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/nam-sang.php"> Sang Nam</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Teaching Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 220</span>
																										<a class="profile__email" href="mailto:snam@uvic.ca">snam@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-721-6402">250-721-6402</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Ignace Ng of the Gustavson School of Business" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/ignace-ng.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/ng-ignace.php"> Ignace Ng</a>
										</p>
										<p>
																		<strong class="profile__title">Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 428</span>
																										<a class="profile__email" href="mailto:ing@uvic.ca">ing@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-721-6073">250-721-6073</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Andrew Park" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/andrew-park.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/park-andrew.php"> Andrew Park</a>
										</p>
										<p>
																		<strong class="profile__title">Assistant Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 412</span>
																										<a class="profile__email" href="mailto:apark1@uvic.ca">apark1@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-721-8209">250-721-8209</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson faculty member Alison Parker" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/alison-parker.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/parker-alison.php"> Alison Parker</a>
										</p>
										<p>
																		<strong class="profile__title">Assistant Teaching Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 424</span>
																										<a class="profile__email" href="mailto:alisonjparker@uvic.ca">alisonjparker@uvic.ca</a>
																						</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Simon Pek" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/simon-pek.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/pek-simon.php"> Simon Pek</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor; Associate Director, Centre for Social and Sustainable Innovation</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 476</span>
																										<a class="profile__email" href="mailto:spek@uvic.ca">spek@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-472-5349">250-472-5349</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Heather Ranson of the Gustavson School of Business" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/heather-ranson.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/ranson-heather.php"> Heather Ranson</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Teaching Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 208</span>
																										<a class="profile__email" href="mailto:hranson@uvic.ca">hranson@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-721-6112">250-721-6112</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Generic person image" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/image-missing.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/reed-jennifer.php"> Jennifer Reed</a>
										</p>
										<p>
																		<strong class="profile__title">Assistant Teaching Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 216</span>
																										<a class="profile__email" href="mailto:jereed@uvic.ca">jereed@uvic.ca</a>
																						</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson faculty member Sorin Rizeanu" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/sorin-rizeanu.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/rizeanu-sorin.php"> Sorin Rizeanu</a>
										</p>
										<p>
																		<strong class="profile__title">Assistant Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 478</span>
																										<a class="profile__email" href="mailto:srizeanu@uvic.ca">srizeanu@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-472-5382">250-472-5382</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Yan Shen" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/yan-shen.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/shen-yan.php"> Yan Shen</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor; Academic Director, MGB Program</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 444</span>
																										<a class="profile__email" href="mailto:yanshen@uvic.ca">yanshen@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-721-6312">250-721-6312</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Linda Shi of the Gustavson School of Business" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/linda-shi.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/shi-linda-hui.php"> Linda Hui Shi</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 210</span>
																										<a class="profile__email" href="mailto:lshi@uvic.ca">lshi@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-721-6408">250-721-6408</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Natalie Slawinski" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/natalie-slawinski.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/slawinski-natalie.php"> Natalie Slawinski</a>
										</p>
										<p>
																		<strong class="profile__title">Professor; Director, Centre for Social and Sustainable Innovation (CSSI)</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 432</span>
																										<a class="profile__email" href="mailto:nslawinski@uvic.ca">nslawinski@uvic.ca</a>
																						</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Claudia Smith of the Gustavson School of Business" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/claudia-smith.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/smith-claudia.php"> Claudia Smith</a>
										</p>
										<p>
																		<strong class="profile__title">Assistant Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 264</span>
																										<a class="profile__email" href="mailto:smithcg@uvic.ca">smithcg@uvic.ca</a>
																						</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson's J. Brock Smith" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/brock-smith.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/smith-j-brock.php"> J. Brock Smith</a>
										</p>
										<p>
																		<strong class="profile__title">Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 232</span>
																										<a class="profile__email" href="mailto:smithb@uvic.ca">smithb@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-721-6070">250-721-6070</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Stuart Snaith of the Gustavson School of Business" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/stuart-snaith.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/snaith-stuart.php"> Stuart Snaith</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 238</span>
																										<a class="profile__email" href="mailto:ssnaith@uvic.ca">ssnaith@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-721-6646">250-721-6646</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson's Doug Stuart" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/doug-stuart-2024.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/stuart-douglas.php"> Douglas Stuart</a>
										</p>
										<p>
																		<strong class="profile__title">Assistant Teaching Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 422</span>
																										<a class="profile__email" href="mailto:dstuart@uvic.ca">dstuart@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-472-4718">250-472-4718</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Kenneth Thornicroft of the Gustavson School of Business" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/kenneth-thornicroft.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/thornicroft-kenneth-wm.php"> Kenneth Wm. Thornicroft</a>
										</p>
										<p>
																		<strong class="profile__title">Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 206</span>
																										<a class="profile__email" href="mailto:kthornic@uvic.ca">kthornic@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-721-6424">250-721-6424</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Christian Van Buskirk from the Gustavson School of Business" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/christian-van-buskirk.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/van-buskirk-christian.php"> Christian Van Buskirk</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Teaching Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 262</span>
																										<a class="profile__email" href="mailto:cvanbus@uvic.ca">cvanbus@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-472-4845">250-472-4845</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Liana Victorino of the Gustavson School of Business" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/liana-victorino.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/victorino-liana.php"> Liana Victorino</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 452</span>
																										<a class="profile__email" href="mailto:lianav@uvic.ca">lianav@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-721-6400">250-721-6400</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Jie Zhang of the Gustavson School of Business" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/jie-zhang.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/zhang-jie.php"> Jie Zhang</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 450</span>
																										<a class="profile__email" href="mailto:jiezhang@uvic.ca">jiezhang@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-472-5736">250-472-5736</a>
														</p>
																		</div>
				</div>
						</div>
																																																																										<div class="list-group-item filtered__item">
								<div class="row">
			<div class="col-sm-4">
																																	<picture>
																																																																				<source media="(max-width: 380px)" srcset="https://www.uvic.ca/assets/core-4-0/img/blank.gif"/>
								<img alt="Gustavson researcher Sarah Zheng" class="profile__img" height="303" loading="lazy" src="../../_assets/images/profiles/sarah-zheng.jpg" width="240"/>
																		</picture>
								</div>
			<div class="col-sm-8">
												<p class="profile__name h4">
														<a href="profiles/zheng-sarah.php"> Sarah Zheng</a>
										</p>
										<p>
																		<strong class="profile__title">Associate Professor</strong>
																						</p>
																		<p>
																		<span class="profile__office">Office: BEC 438</span>
																										<a class="profile__email" href="mailto:szheng@uvic.ca">szheng@uvic.ca</a>
																										<a class="profile__phone" href="tel:250-853-3217">250-853-3217</a>
														</p>
																		</div>
				</div>
						</div>
																</div>
		</div>

																																																</div>
				</div>

					</div>
					<div class="uv-right">

					</div>
				</div>
			</div>
		</main>
		<footer class="ftr">


<div>
		<div class="ftr__lvl1">
				<div class="container">
						<div class="row justify-content-center">
								<div class="ftr__social">
										<span class="h3 me-4" style="margin: 0 0;">Follow us </span>
										<a aria-label="Gustavson Instagram" class="btn--social" href="https://www.instagram.com/GustavsonUVic/"> <i class="fa-brands fa-instagram">​</i> </a>
										<a aria-label="Gustavson Facebook" class="btn--social" href="https://www.facebook.com/GustavsonUVic"> <i class="fab fa-facebook-square">​</i> </a>
										<a aria-label="Gustavson Twitter footer link" class="btn--social" href="https://twitter.com/GustavsonUVic"> <i class="fab fa-twitter-square">​</i> </a>
										<a aria-label="Gustavson YouTube" class="btn--social" href="https://www.youtube.com/gustavsonuvic"> <i class="fa-brands fa-youtube">​</i> </a>
										<a aria-label="Gustavson LinkedIn" class="btn--social" href="https://ca.linkedin.com/showcase/peter-b-gustavson-school-of-business/"><i class="fa-brands fa-linkedin">​</i> </a>
								</div>
						</div>
				</div>
		</div>
		<div class="ftr__lvl3 bg-white">
				<div class="container position-relative">
						<div class="row ">
								<div class="accred">
										<a class="stretched-link" href="../../about/rankings-and-reputation/index.php#ipn-accreditations">Accreditations</a>
										<img alt="Equis accredited logo" src="../../_assets/images/template/equis.png"/>
										<img alt="AACSB accredited logo" src="../../_assets/images/template/aacsb.png"/>
										<img alt="Principles for Responsible Management Education logo" src="../../_assets/images/template/prime.png"/>
										<img alt="CEWIL Canada logo" src="../../_assets/images/template/cewil.png"/>
								</div>
						</div>
				</div>
		</div>
		<div class="ftr__lvl4">
				<div class="container">
						<div class="row justify-content-center">
								<div class="col col-lg-4">
										<div class="row">
												<div class="col-xl">
														<a href="/index.php"><img alt="University of Victoria Logo" class="ftr__logo" src="https://www.uvic.ca/assets/core-4-0/img/shield-horz-white.svg"/></a>
												</div>
												<div class="col-xl">
														<p>
																University of Victoria <br/>
																Peter B. Gustavson School of Business  <br/>
																PO Box 1700 STN CSC <br/>
																Victoria BC  V8W 2Y2 <br/>
																Canada
														</p>
												</div>
										</div>
								</div>
								<div class="col col-lg-1">
										<div class="row">
												<div class="col-xl">
														<ul class="list-unstyled">
																<li>
																		<span aria-label="Gustavson Phone footer link" class="ftr__iconlink text-nowrap"><i aria-hidden="true" class="fa-solid fa-phone"> </i>1-250-472-4139</span>
																</li>
																<li>
																		<a aria-label="Gustavson Email footer link" class="ftr__iconlink" href="mailto:gustavson@uvic.ca"><i aria-hidden="true" class="fa-solid fa-envelope"> </i>gustavson@uvic.ca</a>
																</li>
														</ul>
												</div>
										</div>
								</div>
						</div>
				</div>
		</div>
		<div class="ftr__lvl5">
				<div class="container">
						<div class="row justify-content-center">
								<div class="col col-lg-8">
										<ul class="list-inline">
												<li class="list-inline-item"><a aria-label="Terms of use footer link" href="/info/terms-of-use/index.php">Terms of use</a></li>
												<li class="list-inline-item"><a aria-label="Accessibility footer link" href="/info/accessibility/index.php">Accessibility</a></li>
												<li class="list-inline-item"><a aria-label="Emergency contacts footer link" href="/info/emergency-contacts/index.php">Emergency contacts</a></li>
										</ul>
										<ul class="list-inline">
												<li class="list-inline-item"><a aria-label="© University of Victoria copyright footer link" href="/info/copyright-information/index.php">© University of Victoria</a></li>
												<li class="list-inline-item"><a aria-label="Website feedback footer link" href="/info/website-feedback/index.php">Website feedback</a></li>
										</ul>
								</div>
						</div>
				</div>
		</div>
		<a class="btn btn--ico btn--top" href="#uv-header"><i aria-hidden="true" class="fa-solid fa-arrow-up"></i> <span class="sr-only">Back to top of page</span></a>
</div>

		</footer>


	<div>
<div class="uv-banner collapse show bg--yellow text--black" id="cookies-banner" role="alert">
	<div aria-describedby="cookie-description" aria-label="Cookie Notice" class="container" role="region">
			<p id="cookie-description">This website stores cookies on your computer. These cookies are used to collect information about how you interact with our website and allow us to remember your browser. We use this information to improve and customize your browsing experience, for analytics and metrics about our visitors both on this website and other media, and for marketing purposes. By using this website, you accept and agree to be bound by UVic’s <a href="/info/terms-of-use/index.php">Terms of Use</a> and <a class="icoa icoa--pdf" href="https://www.uvic.ca/universitysecretary/assets/docs/policies/GV0235.pdf">Protection of Privacy Policy</a>. If you do not agree to the above, you must not use this website.</p>
<div class="btn-row"><button aria-controls="cookies-banner" aria-expanded="false" class="btn btn-outline-primary" data-bs-target="#cookies-banner" data-bs-toggle="collapse" data-target="#cookies-banner" data-toggle="collapse" id="cookies-btn">Close</button></div>
	</div>
</div>

</div>


<!-- JS Pack -->


		<script src="/assets/core-4-1/js/scripts.1714067250.js" type="text/javascript"></script>    

		<!-- Loop through and grab all dependencies flags -->


		<script src="/assets/core-4-1/js/animation/aos.1690244651.js" type="text/javascript"></script>    


<!-- Font Awesome -->
<script crossorigin="anonymous" src="https://kit.fontawesome.com/6353bd3065.js"></script>









	</body>
</html>`
