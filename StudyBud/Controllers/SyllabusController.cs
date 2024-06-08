using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using StudyBud.Business;

namespace StudyBud.Controllers
{
    public class SyllabusController : Controller
    {
        public IActionResult Index()
        {
            return View();
        }

        public IActionResult Upload()
        {
            //here we will upload a syllabus to be analyzed
            //call bl to process syllabus and return a confirmation view
            return View();
        }

   
    }
}