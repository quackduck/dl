<!DOCTYPE html>
<html data-adblockkey='MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAL/3/SrV7P8AsTHMFSpPmYbyv2PkACHwmG9Z+1IFZq3vA54IN7pQcGnhgNo+8SN9r/KtUWCb9OPqTfWM1N4w/EUCAwEAAQ==_mV7PA34SCxFFygw515oL4l4mRvX4gtyH1tUUp7b5JVaCgvz3KsedwT0y5f3hpOLGZcoOq3d9JBns3Sk4/ZOpXw=='>
<head>


<meta charset='UTF-8'>
<title>Foo.com</title>
<link href='data:image/png;base64,iVBORw0KGgo=' rel='icon' type='image/png'>
<link href="/assets/style-85cdb3be1709c9028b204b38f8ccb358.css" media="none" onload="if(media!==&#x27;all&#x27;)media=&#x27;all&#x27;" rel="stylesheet" type="text/css" />
<link href="/assets/belt_layout_caf-4a3dea87ed4a809b49d0640b48d26654.css" media="none" onload="if(media!==&#x27;all&#x27;)media=&#x27;all&#x27;" rel="stylesheet" type="text/css" />
<script src="///ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js" type="text/javascript"></script>
<!-- = stylesheet_link_tag    "/load_style", :media => "all" -->
<meta content="authenticity_token" name="csrf-param" />
<meta content="TR6WRl+H6zsWjBVsfflVVPR5jN4GuvikzhUGeC7BsDI=" name="csrf-token" />
<style>
a, .ourterm a{
	color: #00c;
}

.background1{
	background-color: #ccc;
}

.background2{
	background-color: #aaa;
}

.divider_background{
	background-color: #aaa;
}


</style>

</head>
<body>
<script async="async" src="/assets/application-07ff4bc4011765f42975871a196fae13.js" type="text/javascript"></script>




<center>
<div id='outer_container'>
<div id='outer_container2'>
<div id='container'>
<div id='header'>
<div class='logo'>
<a href="/">Foo.com</a>
</div>
<div class='googleframe' id='header_image'></div>
</div>







<div id='main'>
<div class='image'>
<img alt="Cordovabeach" border="0" src="/media/W1siZiIsIjIwMTIvMDQvMjYvMjAvMTEvNDkvNDI2L2NvcmRvdmFiZWFjaC5qcGciXSxbInAiLCJ0aHVtYiIsIjc1MHgyMDAjIl1d/cordovabeach.jpg" /></div>
<div class='search'>
<div class='catch_phrase'>
<div id='wysiwyg_id_322' class='mercury-region' data-type='editable'>Search FOO.com
</div>
</div>
<div class='googleframe' id='search'></div>
</div>
</div>
<script>
var searchblock1 = {
  'type' : 'searchbox',
  'container' : 'search',
  'widthSearchButton' : 100,
  'searchBoxMethod' : 'get'
};
</script>

<div id='footer'>
<a href='/digimedia_privacy_policy.html' target='_blank'>Privacy Policy</a>
 - 
<a href='http://www.digimedia.com' target='_blank'>Copyright &copy; 2021 Digimedia.com, L.P.</a>
</div>

</div>

</div>
</div>
</center>
<script src="/assets/abp2-4831d5b24977f8140fd9aa25543527f2.js" type="text/javascript"></script>
<script src="/assets/ads-832e97bfd2d3e735f6dc8a30dd7190bc.js" type="text/javascript"></script>
<!-- %script{:type=>"text/javascript", :language=>"JavaScript", :src=>"/abp.js"} -->
<script language='JavaScript' src='///www.google.com/adsense/domains/caf.js' type='text/javascript'></script>
<script>
  function jscript_log(name, severity, message, page){
    $.ajax({
      type: "POST",
      url: "/log_error",
      data: {
        name: name,
        severity: severity,
        domain: window.location.hostname,
        message: message,
        tag: 'dp-digimedia_js',
        page: page
      },
      context: document.body
    })
  }
</script>
<script>
function google_index_loaded(requestAccepted, status) {
  if(!requestAccepted){
    console.log('failed: ' + requestAccepted);
    jscript_log('faillisted', 10, JSON.stringify(status), "/")
  }
  console.log('status: ');
  console.log(status);
}
var pageOptions = {
'pubId' : 'partner-dp-digimedia_js',
'domainRegistrant': 'as-drid-oo-1626960400946279',

'resultsPageBaseUrl': '///www.foo.com/results?',
'channel': 'digi-caf_pef,digimedia-template-09',
'adtest' : false,
'optimizeTerms': true,
'terms': '',



'numRepeated' : 3,
'styleId': '9039920606',
'colorTitleLink' : '#00c',
'pageLoadedCallback' : google_index_loaded
}
</script>

<script>new google.ads.domains.Caf(pageOptions, searchblock1);AdblockPlus.detect('/px.gif', function(usesABP){console.log('usesABP:' + usesABP);if(usesABP){blocked_ads_logger('whitelisted');}else{if(adsblocked()){blocked_ads_logger('blocked');}else{blocked_ads_logger('not_blocked');}}});</script>

<!-- Global site tag (gtag.js) - Google Analytics -->
<script async src='https://www.googletagmanager.com/gtag/js?id=UA-1726084-83'></script>
<script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());
  
  gtag('config', 'UA-1726084-83', {
    'page_title': 'foo.com',
    'page_location': 'http://www.foo.com/',
    'page_path': '/'
  });
  
  
  gtag('event', 'index', {'event_category': 'domain_name', 'event_label': 'foo.com'});
</script>

<script>
var privacy_policy='/digimedia_privacy_policy.html';
var domain_name='foo.com'
</script>
<script src="///privacy.digimedia.com/check_cookie_country_code.js" type="text/javascript"></script>
<style>
  .privacy_consent{
      background:rgb(0,0,0,0.5);
  }
  .privacy_consent a{
    color:black;
  }
  .privacy_consent_box{
    text-shadow:none;
    padding:0px 20px;
    font-size:12pt;
  }
  .privacy_consent_box h2{
    display:none;
    margin:0px;
    padding:15px 0px 5px 0px;
    border-bottom: 1px solid #eee;
  }
  .privacy_consent_button{
    padding:10px 20px;
    background:green;
    color:white;
  }
  .privacy_consent_more_button{
    margin-right:15px;
    padding:10px 20px;
    background:black;
    color:white;
  }
  .privacy_consent_footer{
    position:relative;
    clear:both;
  }
  .privacy_consent_footer_p{
    font-size:0.6em;
    margin:0px;
    margin-top:15px;
  }
</style>
<script>
  $(document).ready(function(){
    pc_options = {
      privacy_domain: "privacy.digimedia.com",
      ssl: false,
      title: null,
      question: "This website uses cookies to improve user experience. By using our website you consent to our <a href='"+privacy_policy+"' target='_blank'>privacy policy</a>.",
      links_left: [],
      agree_text: "YES, I ACCEPT",
      more_info_text: "MORE INFO",
      more_info_link: privacy_policy,
      text_color: 'black',
      link_color: 'black',
      width: '320px',
      height: '150px'
    }
    privacy_consent(pc_options);
  });
</script>

</body>
</html>
