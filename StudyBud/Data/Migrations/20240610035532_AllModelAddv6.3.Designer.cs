﻿// <auto-generated />
using System;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Infrastructure;
using Microsoft.EntityFrameworkCore.Migrations;
using Microsoft.EntityFrameworkCore.Storage.ValueConversion;
using StudyBud.Data;

#nullable disable

namespace StudyBud.Data.Migrations
{
    [DbContext(typeof(ApplicationDbContext))]
    [Migration("20240610035532_AllModelAddv6.3")]
    partial class AllModelAddv63
    {
        /// <inheritdoc />
        protected override void BuildTargetModel(ModelBuilder modelBuilder)
        {
#pragma warning disable 612, 618
            modelBuilder.HasAnnotation("ProductVersion", "7.0.19");

            modelBuilder.Entity("Microsoft.AspNetCore.Identity.IdentityRole", b =>
                {
                    b.Property<string>("Id")
                        .HasColumnType("TEXT");

                    b.Property<string>("ConcurrencyStamp")
                        .IsConcurrencyToken()
                        .HasColumnType("TEXT");

                    b.Property<string>("Name")
                        .HasMaxLength(256)
                        .HasColumnType("TEXT");

                    b.Property<string>("NormalizedName")
                        .HasMaxLength(256)
                        .HasColumnType("TEXT");

                    b.HasKey("Id");

                    b.HasIndex("NormalizedName")
                        .IsUnique()
                        .HasDatabaseName("RoleNameIndex");

                    b.ToTable("AspNetRoles", (string)null);
                });

            modelBuilder.Entity("Microsoft.AspNetCore.Identity.IdentityRoleClaim<string>", b =>
                {
                    b.Property<int>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("INTEGER");

                    b.Property<string>("ClaimType")
                        .HasColumnType("TEXT");

                    b.Property<string>("ClaimValue")
                        .HasColumnType("TEXT");

                    b.Property<string>("RoleId")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.HasKey("Id");

                    b.HasIndex("RoleId");

                    b.ToTable("AspNetRoleClaims", (string)null);
                });

            modelBuilder.Entity("Microsoft.AspNetCore.Identity.IdentityUserClaim<string>", b =>
                {
                    b.Property<int>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("INTEGER");

                    b.Property<string>("ClaimType")
                        .HasColumnType("TEXT");

                    b.Property<string>("ClaimValue")
                        .HasColumnType("TEXT");

                    b.Property<string>("UserId")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.HasKey("Id");

                    b.HasIndex("UserId");

                    b.ToTable("AspNetUserClaims", (string)null);
                });

            modelBuilder.Entity("Microsoft.AspNetCore.Identity.IdentityUserLogin<string>", b =>
                {
                    b.Property<string>("LoginProvider")
                        .HasMaxLength(128)
                        .HasColumnType("TEXT");

                    b.Property<string>("ProviderKey")
                        .HasMaxLength(128)
                        .HasColumnType("TEXT");

                    b.Property<string>("ProviderDisplayName")
                        .HasColumnType("TEXT");

                    b.Property<string>("UserId")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.HasKey("LoginProvider", "ProviderKey");

                    b.HasIndex("UserId");

                    b.ToTable("AspNetUserLogins", (string)null);
                });

            modelBuilder.Entity("Microsoft.AspNetCore.Identity.IdentityUserRole<string>", b =>
                {
                    b.Property<string>("UserId")
                        .HasColumnType("TEXT");

                    b.Property<string>("RoleId")
                        .HasColumnType("TEXT");

                    b.HasKey("UserId", "RoleId");

                    b.HasIndex("RoleId");

                    b.ToTable("AspNetUserRoles", (string)null);
                });

            modelBuilder.Entity("Microsoft.AspNetCore.Identity.IdentityUserToken<string>", b =>
                {
                    b.Property<string>("UserId")
                        .HasColumnType("TEXT");

                    b.Property<string>("LoginProvider")
                        .HasMaxLength(128)
                        .HasColumnType("TEXT");

                    b.Property<string>("Name")
                        .HasMaxLength(128)
                        .HasColumnType("TEXT");

                    b.Property<string>("Value")
                        .HasColumnType("TEXT");

                    b.HasKey("UserId", "LoginProvider", "Name");

                    b.ToTable("AspNetUserTokens", (string)null);
                });

            modelBuilder.Entity("StudyBud.Models.Assignment", b =>
                {
                    b.Property<string>("AssignmentId")
                        .HasColumnType("TEXT");

                    b.Property<string>("Description")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<string>("Name")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<decimal>("PercentOfGrade")
                        .HasColumnType("TEXT");

                    b.Property<string>("SyllabusID")
                        .HasColumnType("TEXT");

                    b.Property<string>("Type")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.HasKey("AssignmentId");

                    b.HasIndex("SyllabusID");

                    b.ToTable("Assignments");
                });

            modelBuilder.Entity("StudyBud.Models.Book", b =>
                {
                    b.Property<string>("BookId")
                        .HasColumnType("TEXT");

                    b.Property<string>("Author")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<string>("Description")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<string>("ISBN")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<int>("Length")
                        .HasColumnType("INTEGER");

                    b.Property<string>("SyllabusID")
                        .HasColumnType("TEXT");

                    b.Property<string>("Title")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.HasKey("BookId");

                    b.HasIndex("SyllabusID");

                    b.ToTable("Books");
                });

            modelBuilder.Entity("StudyBud.Models.Cohort", b =>
                {
                    b.Property<string>("CohortId")
                        .HasColumnType("TEXT");

                    b.Property<bool>("Mentor")
                        .HasColumnType("INTEGER");

                    b.Property<bool>("StandardUser")
                        .HasColumnType("INTEGER");

                    b.Property<string>("Topic")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<string>("UserId")
                        .HasColumnType("TEXT");

                    b.HasKey("CohortId");

                    b.HasIndex("UserId");

                    b.ToTable("Cohorts");
                });

            modelBuilder.Entity("StudyBud.Models.Degree", b =>
                {
                    b.Property<string>("DegreeId")
                        .HasColumnType("TEXT");

                    b.Property<string>("DegreeType")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<float>("GPA")
                        .HasColumnType("REAL");

                    b.Property<bool>("Graduated")
                        .HasColumnType("INTEGER");

                    b.Property<string>("IssuingSchoolSchoolId")
                        .HasColumnType("TEXT");

                    b.Property<string>("UserId")
                        .HasColumnType("TEXT");

                    b.Property<DateTime>("YearFinished")
                        .HasColumnType("TEXT");

                    b.Property<DateTime>("YearStarted")
                        .HasColumnType("TEXT");

                    b.HasKey("DegreeId");

                    b.HasIndex("IssuingSchoolSchoolId");

                    b.HasIndex("UserId");

                    b.ToTable("Degrees");
                });

            modelBuilder.Entity("StudyBud.Models.Professor", b =>
                {
                    b.Property<string>("ProfessorId")
                        .HasColumnType("TEXT");

                    b.Property<string>("Email")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<string>("Name")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<string>("Phone")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<string>("School")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.HasKey("ProfessorId");

                    b.ToTable("Professors");
                });

            modelBuilder.Entity("StudyBud.Models.School", b =>
                {
                    b.Property<string>("SchoolId")
                        .HasColumnType("TEXT");

                    b.Property<string>("Address")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<bool>("CurrentlyEnrolled")
                        .HasColumnType("INTEGER");

                    b.Property<string>("DegreeTypeInProgress")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<bool>("Hybrid")
                        .HasColumnType("INTEGER");

                    b.Property<bool>("InPerson")
                        .HasColumnType("INTEGER");

                    b.Property<string>("Name")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<bool>("Online")
                        .HasColumnType("INTEGER");

                    b.Property<string>("UserId")
                        .HasColumnType("TEXT");

                    b.Property<int>("Year")
                        .HasColumnType("INTEGER");

                    b.HasKey("SchoolId");

                    b.HasIndex("UserId");

                    b.ToTable("Schools");
                });

            modelBuilder.Entity("StudyBud.Models.Syllabus", b =>
                {
                    b.Property<string>("SyllabusID")
                        .HasColumnType("TEXT");

                    b.Property<string>("ClassTitle")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<byte[]>("Content")
                        .IsRequired()
                        .HasColumnType("BLOB");

                    b.Property<string>("CourseObjectives")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<int>("CreditHours")
                        .HasColumnType("INTEGER");

                    b.Property<DateTime>("EndDate")
                        .HasColumnType("TEXT");

                    b.Property<bool>("LateWork")
                        .HasColumnType("INTEGER");

                    b.Property<string>("Misc")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<string>("Objectives")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<string>("ProfessorId")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<string>("School")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<string>("Semester")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<DateTime>("StartDate")
                        .HasColumnType("TEXT");

                    b.Property<bool>("TA")
                        .HasColumnType("INTEGER");

                    b.Property<string>("TechRequirements")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<string>("UserId")
                        .HasColumnType("TEXT");

                    b.HasKey("SyllabusID");

                    b.HasIndex("ProfessorId");

                    b.HasIndex("UserId");

                    b.ToTable("Syllabi");
                });

            modelBuilder.Entity("StudyBud.Models.User", b =>
                {
                    b.Property<string>("Id")
                        .HasColumnType("TEXT");

                    b.Property<int>("AccessFailedCount")
                        .HasColumnType("INTEGER");

                    b.Property<string>("Address")
                        .HasColumnType("TEXT");

                    b.Property<string>("ConcurrencyStamp")
                        .IsConcurrencyToken()
                        .HasColumnType("TEXT");

                    b.Property<string>("Email")
                        .HasMaxLength(256)
                        .HasColumnType("TEXT");

                    b.Property<bool>("EmailConfirmed")
                        .HasColumnType("INTEGER");

                    b.Property<string>("FName")
                        .HasColumnType("TEXT");

                    b.Property<float?>("GPA")
                        .HasColumnType("REAL");

                    b.Property<string>("LName")
                        .HasColumnType("TEXT");

                    b.Property<bool>("LockoutEnabled")
                        .HasColumnType("INTEGER");

                    b.Property<DateTimeOffset?>("LockoutEnd")
                        .HasColumnType("TEXT");

                    b.Property<string>("MInitial")
                        .HasColumnType("TEXT");

                    b.Property<string>("Name")
                        .HasColumnType("TEXT");

                    b.Property<string>("NormalizedEmail")
                        .HasMaxLength(256)
                        .HasColumnType("TEXT");

                    b.Property<string>("NormalizedUserName")
                        .HasMaxLength(256)
                        .HasColumnType("TEXT");

                    b.Property<string>("PasswordHash")
                        .HasColumnType("TEXT");

                    b.Property<string>("PhoneNumber")
                        .HasColumnType("TEXT");

                    b.Property<bool>("PhoneNumberConfirmed")
                        .HasColumnType("INTEGER");

                    b.Property<string>("SecurityStamp")
                        .HasColumnType("TEXT");

                    b.Property<bool?>("Subscribed")
                        .HasColumnType("INTEGER");

                    b.Property<bool>("TwoFactorEnabled")
                        .HasColumnType("INTEGER");

                    b.Property<string>("UserName")
                        .HasMaxLength(256)
                        .HasColumnType("TEXT");

                    b.HasKey("Id");

                    b.HasIndex("NormalizedEmail")
                        .HasDatabaseName("EmailIndex");

                    b.HasIndex("NormalizedUserName")
                        .IsUnique()
                        .HasDatabaseName("UserNameIndex");

                    b.ToTable("AspNetUsers", (string)null);
                });

            modelBuilder.Entity("Microsoft.AspNetCore.Identity.IdentityRoleClaim<string>", b =>
                {
                    b.HasOne("Microsoft.AspNetCore.Identity.IdentityRole", null)
                        .WithMany()
                        .HasForeignKey("RoleId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();
                });

            modelBuilder.Entity("Microsoft.AspNetCore.Identity.IdentityUserClaim<string>", b =>
                {
                    b.HasOne("StudyBud.Models.User", null)
                        .WithMany()
                        .HasForeignKey("UserId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();
                });

            modelBuilder.Entity("Microsoft.AspNetCore.Identity.IdentityUserLogin<string>", b =>
                {
                    b.HasOne("StudyBud.Models.User", null)
                        .WithMany()
                        .HasForeignKey("UserId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();
                });

            modelBuilder.Entity("Microsoft.AspNetCore.Identity.IdentityUserRole<string>", b =>
                {
                    b.HasOne("Microsoft.AspNetCore.Identity.IdentityRole", null)
                        .WithMany()
                        .HasForeignKey("RoleId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.HasOne("StudyBud.Models.User", null)
                        .WithMany()
                        .HasForeignKey("UserId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();
                });

            modelBuilder.Entity("Microsoft.AspNetCore.Identity.IdentityUserToken<string>", b =>
                {
                    b.HasOne("StudyBud.Models.User", null)
                        .WithMany()
                        .HasForeignKey("UserId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();
                });

            modelBuilder.Entity("StudyBud.Models.Assignment", b =>
                {
                    b.HasOne("StudyBud.Models.Syllabus", null)
                        .WithMany("Assignments")
                        .HasForeignKey("SyllabusID");
                });

            modelBuilder.Entity("StudyBud.Models.Book", b =>
                {
                    b.HasOne("StudyBud.Models.Syllabus", null)
                        .WithMany("Books")
                        .HasForeignKey("SyllabusID");
                });

            modelBuilder.Entity("StudyBud.Models.Cohort", b =>
                {
                    b.HasOne("StudyBud.Models.User", null)
                        .WithMany("Cohorts")
                        .HasForeignKey("UserId");
                });

            modelBuilder.Entity("StudyBud.Models.Degree", b =>
                {
                    b.HasOne("StudyBud.Models.School", "IssuingSchool")
                        .WithMany()
                        .HasForeignKey("IssuingSchoolSchoolId");

                    b.HasOne("StudyBud.Models.User", null)
                        .WithMany("Degrees")
                        .HasForeignKey("UserId");

                    b.Navigation("IssuingSchool");
                });

            modelBuilder.Entity("StudyBud.Models.School", b =>
                {
                    b.HasOne("StudyBud.Models.User", null)
                        .WithMany("Schools")
                        .HasForeignKey("UserId");
                });

            modelBuilder.Entity("StudyBud.Models.Syllabus", b =>
                {
                    b.HasOne("StudyBud.Models.Professor", "Professor")
                        .WithMany()
                        .HasForeignKey("ProfessorId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.HasOne("StudyBud.Models.User", null)
                        .WithMany("Syllabi")
                        .HasForeignKey("UserId");

                    b.Navigation("Professor");
                });

            modelBuilder.Entity("StudyBud.Models.Syllabus", b =>
                {
                    b.Navigation("Assignments");

                    b.Navigation("Books");
                });

            modelBuilder.Entity("StudyBud.Models.User", b =>
                {
                    b.Navigation("Cohorts");

                    b.Navigation("Degrees");

                    b.Navigation("Schools");

                    b.Navigation("Syllabi");
                });
#pragma warning restore 612, 618
        }
    }
}