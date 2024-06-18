using System.Diagnostics;
using Microsoft.AspNetCore.Mvc;
using StudyBud.Models;
using Microsoft.AspNetCore.Identity;

namespace StudyBud.Controllers;

public class HomeController : Controller
{
    private readonly ILogger<HomeController> _logger;
    private readonly SignInManager<User> _signInManager;

    public HomeController(ILogger<HomeController> logger, SignInManager<User> signInManager)
    {
        _logger = logger;
        _signInManager = signInManager;
    }

    public IActionResult Index()
    {
        if (_signInManager.IsSignedIn(User))
        {
            return View();
        }

        return RedirectToPage("/Account/Login", new { area = "Identity"});
    }

    public IActionResult Privacy()
    {
        return View();
    }

    [ResponseCache(Duration = 0, Location = ResponseCacheLocation.None, NoStore = true)]
    public IActionResult Error()
    {
        return View(new ErrorViewModel { RequestId = Activity.Current?.Id ?? HttpContext.TraceIdentifier });
    }
}

