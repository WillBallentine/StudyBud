using System;
using System.Collections.Generic;
using System.IO.Compression;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using StudyBud.Business;
using StudyBud.Data;
using StudyBud.Data.Interfaces;
using StudyBud.Models;

namespace StudyBud.Controllers
{
    public class SyllabusController : Controller
    {
        ISyllabusDAL _syllabusDal;

        public SyllabusController(ISyllabusDAL syllabusDal)
        {
            _syllabusDal = syllabusDal;
        }
        
        public IActionResult Index()
        {
            return View();
        }


    }
}